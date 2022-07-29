package main

import (
	"fmt"
	"time"
)

//往终端写日志相关内容

type Logger struct {
	level LogLevel
}

//构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		level: level,
	}
}

//日志开关：大于debug的才输出
func (l Logger) enable(logLevel LogLevel) bool {
	return logLevel >= l.level
}

func log(lv LogLevel, format string,a ...interface{}) {
	msg := fmt.Sprintf(format,a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(2)
	fmt.Printf("[%s] [%s] [%s-%s-%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv),funcName, fileName, lineNo, msg)
}

func (l Logger) Debug(format string,a ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format,a...)
	}
}
func (l Logger) Info(format string,a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format,a...)
	}
}
func (l Logger) Warning(format string,a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format,a...)
	}
}
func (l Logger) Error(format string,a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format,a...)
	}
}
func (l Logger) Fatal(format string,a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL,format,a...)
	}
}
