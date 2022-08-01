package main

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//定义日志级别
type LogLevel uint16

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}
const (
	UNKOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//根据传递进来的级别找到对应的类型
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKOWN, err
	}
}

//根据类型找到对应级别
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

//拿到调用函数的具体信息
func getInfo(skip int)(funcName,fileName string,lineNo int) {
	pc,file,lineNo,ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName,".")[1]
	return
}


