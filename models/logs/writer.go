// Copyright 2018 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logs

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	ignoreFunc = []string{
		"views.RenderHandlerError",
		"views.RenderInternalError",
		"RenderHandlerError",
		"RenderInternalError",
		"ErrorStack",
		"ErrorLogHandler",
		"v1.Catch",
		"runtime.gopanic",
		"runtime.panicindex",
		"runtime.call32",
		"runtime.panicdottypeE",
		"v1.WrapAPIHandler.func1",
	}
	ignoreFiles = []string{
		"asm_amd64",
		"asm_amd64.s",
		"server.go",
		"testing.go",
	}
)

// Fatal - output formated (function and line calls) fatal information
func Fatal(err error, args ...interface{}) {
	pc, _, _, _ := runtime.Caller(2)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logErr.Output(2, fmt.Sprintf("[FATAL];%v;%v;%v", changeShortName(runtime.FuncForPC(pc).Name()),
		err, getArgsString(args...)))
	os.Exit(1)

}
func changeShortName(file string) (short string) {
	short = file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}

// DebugLog output formated(function and line calls) debug information
func DebugLog(args ...interface{}) {
	if *fDebug {
		pc, _, _, _ := runtime.Caller(logDebug.calldepth - 2)
		logDebug.funcName = changeShortName(runtime.FuncForPC(pc).Name())

		logDebug.Printf(args...)
	}
}

// DebugLog output formated(function and line calls) debug information
func TraceLog(args ...interface{}) {
	if *fDebug {
		logDebug.Printf(args...)
	}
}

// StatusLog output formated information for status
func StatusLog(args ...interface{}) {
	if *fStatus {
		logStat.Printf(args...)
	}
}

// ErrorLog - output formated(function and line calls) error information
func ErrorLog(err error, args ...interface{}) {
	calldepth := logErr.calldepth

	isIgnore := true

	for pc, _, _, ok := runtime.Caller(calldepth); ok && isIgnore; pc, _, _, ok = runtime.Caller(calldepth) {
		calldepth++
		logErr.funcName = changeShortName(runtime.FuncForPC(pc).Name())
		// пропускаем рендер ошибок
		isIgnore = isIgnoreFunc(logErr.funcName)
	}
	//todo: print standart
	var message string
	if err != nil {
		message = fmt.Sprintf("%s;%s;%s", logErr.funcName,
			err.Error(), getArgsString(args...))
	} else {
		message = fmt.Sprintf("%s;%s", logErr.funcName,
			getArgsString(args...))

	}

	logErr.Output(calldepth, message)

}

// ErrorStack - output formatted(function and line calls) error runtime stack information
func ErrorStack(err error, args ...interface{}) {

	i := stackBeginWith
	mes := fmt.Sprintf("[ERROR_STACK];%s;%s ", err, getArgsString(args...))

	for pc, file, line, ok := runtime.Caller(i); ok; pc, file, line, ok = runtime.Caller(i) {
		i++
		funcName := changeShortName(runtime.FuncForPC(pc).Name())
		// пропускаем рендер ошибок
		isIgnore := isIgnoreFunc(funcName)
		if !isIgnore {
			runFile := changeShortName(file)
			isIgnore = isIgnoreFile(runFile)
			if !isIgnore {
				logErr.Output(i, mes+fmt.Sprintf("%s:%d: %s; ", runFile, line, funcName))
			}
		}
	}
}

func isIgnoreFile(runFile string) bool {
	for _, name := range ignoreFiles {
		if (runFile == name) || (strings.HasPrefix(runFile, name)) {
			return true
		}
	}
	return false
}

func isIgnoreFunc(funcName string) bool {
	for _, name := range ignoreFunc {
		if (funcName == name) || (strings.HasSuffix(funcName, "."+name)) {
			return true
		}
	}
	return false
}

// ErrorLogHandler - output formated(function and line calls) error information
func ErrorLogHandler(err error, args ...interface{}) {
	ErrorStack(err, args...)

}
