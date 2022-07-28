package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//往终端写日志相关内容

//定义日志级别
type LogLevel uint16

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
	return  logLevel >= l.level
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [INFO] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [WARNING] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [ERROR] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [FATAL] %s\n",now.Format("2006-01-02 15:04:05"),msg)
	}
}
