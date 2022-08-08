package main

import (
	"fmt"
	"time"
)

//往终端写日志相关内容

type ConsleLogger struct {
	level LogLevel
}

//构造函数
func NewLog(levelStr string) ConsleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsleLogger{
		level: level,
	}
}

//日志开关：大于debug的才输出
func (l ConsleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= l.level
}

func (c ConsleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [%s] [%s-%s-%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

func (c ConsleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}
func (c ConsleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}
func (c ConsleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}
func (c ConsleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
func (c ConsleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
