// Copyright 2018 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package logs output logs and advanced debug information
package logs

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

var (
	fDebug   = flag.Bool("debug", false, "debug mode")
	fStatus  = flag.Bool("status", true, "status mode")
	logErr   = NewWrapKitLogger(colors[ERROR]+"ERROR"+LogEndColor, 1)
	logStat  = NewWrapKitLogger("INFO", 3)
	logDebug = NewWrapKitLogger(colors[DEBUG]+"DEBUG"+LogEndColor, 3)
)

// LogsType - interface for print logs record
type LogsType interface {
	PrintToLogs() string
}

type wrapKitLogger struct {
	*log.Logger
	calldepth int
	funcName  string
	typeLog   string
	toSentry  bool
	toOther   io.Writer
}

const logFlags = log.Lshortfile | log.Ltime

var stackBeginWith = 1

func NewWrapKitLogger(pref string, depth int) *wrapKitLogger {
	return &wrapKitLogger{
		Logger:    log.New(os.Stdout, "[["+pref+"]]", logFlags),
		typeLog:   pref,
		calldepth: depth,
	}
}

// SetDebug set debug level for log, return old value
func SetDebug(d bool) bool {
	old := *fDebug
	*fDebug = d

	return old
}

// SetStatus set status level for log, return old value
func SetStatus(s bool) bool {
	old := *fStatus
	*fStatus = s

	return old
}

type FgLogWriter int8

const (
	FgAll FgLogWriter = iota
	FgErr
	FgInfo
	FgDebug
)

func (logger *wrapKitLogger) addWriters(newWriter io.Writer) {
	if logger.toOther == nil {
		logger.toOther = newWriter
	} else {
		logger.toOther = io.MultiWriter(logger.toOther, newWriter)
	}
}

// SetWriters for logs
func SetWriters(newWriter io.Writer, logFlag FgLogWriter) {

	switch logFlag {
	case FgAll:
		logErr.addWriters(newWriter)
		logStat.addWriters(newWriter)
		logDebug.addWriters(newWriter)
	case FgErr:
		logErr.addWriters(newWriter)
	case FgInfo:
		logStat.addWriters(newWriter)
	case FgDebug:
		logDebug.addWriters(newWriter)
	}
}

// SetSentry set SetSentry output for error
func SetSentry(dns string) error {
	err := sentry.Init(sentry.ClientOptions{Dsn: dns})
	if err != nil {
		return errors.Wrap(err, "sentry.Init")
	}

	logErr.toSentry = true

	return nil
}

// SetLogFlags set logger flags & return old flags
func SetLogFlags(f int) int {

	flags := logErr.Flags()

	logErr.SetFlags(f)
	logDebug.SetFlags(f)
	logStat.SetFlags(f)

	return flags
}

// SetStackBeginWith set stackBeginWith level for log, return old value
func SetStackBeginWith(s int) int {
	old := stackBeginWith
	stackBeginWith = s

	return old
}

type logMess struct {
	Message string    `json:"message"`
	Now     time.Time `json:"@timestamp"`
	Level   string    `json:"level"`
	//vars    [] interface{} `json:"vars"`
}

func NewlogMess(mess string, logger *wrapKitLogger) *logMess {
	return &logMess{mess, time.Now(), logger.typeLog}
}

func (logger *wrapKitLogger) Printf(vars ...interface{}) {

	checkPrint, checkType := vars[0].(errLogPrint)

	if checkType == true {
		vars = vars[1:]
	}

	mess := getArgsString(vars...)

	// changed for TestErrorLogTelegramWrite() because localhost 
	// doesn't hear the request with "\n" at the end
	if logger.toOther != nil {
		go logger.toOther.Write([]byte(mess))
	}	

	mess += "\n"
	if checkType && (checkPrint == true) {
		fmt.Printf(mess)
	} else {
		logger.Output(logger.calldepth, mess)
	}

}

func getArgsString(args ...interface{}) (message string) {

	if len(args) < 1 {
		return ""
	}
	// if first param is formating string
	if format, ok := args[0].(string); ok && (strings.Index(format, "%") > -1) {
		return fmt.Sprintf(format, args[1:]...)
	}

	comma := ""
	for _, arg := range args {

		switch val := arg.(type) {
		case nil:
			message += comma + " is nil"
		case string:
			message += comma + val
		case []string:
			for _, value := range val {
				message += value + "\n"
			}
		case LogsType:
			message += comma + val.PrintToLogs()
		case time.Time:
			message += comma + val.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
		case []interface{}:
			if len(val) > 0 {
				message += comma + getArgsString(val...)
			}
		case error:
			message += comma + fmt.Sprintf("%v", arg)
		default:

			message += comma + fmt.Sprintf("%#v", arg)
		}

		comma = ", "
	}

	return message
}
