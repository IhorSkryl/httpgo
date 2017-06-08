// Copyright 2017 Author: Sergey Litvinov. All rights reserved.
//@package logs Output logs and advanced debug information
//ErrorLog(err error, args ...interface{}) - output formated(function and line calls) error information
//ErrorStack() - output formated(function and line calls) error runtime stack information
//FatalLog(err error, args ...interface{}) - output formated (function and line calls) fatal information
//StatusLog(err error, args ...interface{}) - output formated information for status
package logs

import (
    "runtime"
    "log"
    "flag"
	"fmt"
)

var F_debug    = flag.String("debug","","debug mode")
var F_status   = flag.String("status"," ","status mode")

//DebugLog( args ...interface{}) - output formated(function and line calls) debug information
//@version 1.1 2017-05-31 Sergey Litvinov - Remote requred args
func DebugLog( args ...interface{}) {
    _, fn, line, _ := runtime.Caller(1)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
    if *F_debug > "" {
	    for _, arg := range args {

		    log.Output(2, fmt.Sprintf("[DEBUG];%s;in line;%d;%v", fn, line, arg))
	    }
    }
}

//StatusLog(err error, args ...interface{}) - output formated information for status
//@version 1.0 2017-05-31 Sergey Litvinov - Create
func StatusLog( args ...interface{}) {
	//_, fn, line, _ := runtime.Caller(1)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if *F_status > "" {
		log.Output(2, fmt.Sprintf("[STATUS];;;;%v",  args) )
	}
}

//ErrorLog(err error, args ...interface{}) - output formated(function and line calls) error information
//@version 1.1 2017-05-31 Sergey Litvinov - Remote requred advanced arg
func ErrorLog(err error, args ...interface{}) {
    pc, _, _, _ := runtime.Caller(1)

	errLog := log.Output(2, fmt.Sprintf("[ERROR];%s %v;%v", changeShortName(runtime.FuncForPC(pc).Name()), err, args) )
	if errLog != nil {
		log.Print(errLog)
	}
}


//ErrorStack() - output formated(function and line calls) error runtime stack information
//@version 1.00 2017-06-02 Sergey Litvinov - Create
func ErrorStack() {
	i:=0
	for {
		pc, fn, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		log.Output( i+1, fmt.Sprintf("[ERROR_STACK];%s[%s:%d];;", changeShortName(runtime.FuncForPC(pc).Name()), changeShortName(fn), line ) )
		i++
	}
}

//FatalLog(err error, args ...interface{}) - output formated (function and line calls) fatal information
//@version 1.0 2017-05-31 Sergey Litvinov - Create
func Fatal(err error, args ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Fatalf("[FATAL];%s;in line;%d;%v;%v", fn, line, err, args)
}

func changeShortName(file string)(short string){
	short = file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	//file1 = short
	return short
}
