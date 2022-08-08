package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志

type FileLogger struct {
	Level       LogLevel
	filePath    string //保存路径
	fileName    string //保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 //切割的大小
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

//根据指定的日志文件路径和名字打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err file failed,err:%v\n", err)
		return err
	}
	//日志文件都打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

//日志开关：大于debug的才输出
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

//判断文件是否需要切割
func (f *FileLogger) checkFile(file *os.File) bool {
	fileInfo,err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n",err)
		return false
	}
	//如果当前文件大小 大于等于 日志文件的最大值，返回true
	return fileInfo.Size() >= f.maxFileSize
}

//切割
func (f *FileLogger) splitFile(file *os.File)(*os.File,error) {
	nowStr := time.Now().Format("20060102150405000")
	fileInof,err := file.Stat()
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n",err)
		return nil,err
	}
	logName := path.Join(f.filePath,fileInof.Name())  //当前的日志文件路径
	newLogName := fmt.Sprintf("%s.bak%s",logName,nowStr) //拼接一个日志文件备份的名字
	//切割日志文件
	//1、关闭当前的文件
	file.Close()
	//2、备份 rename  在后面加上时间
	os.Rename(logName,newLogName)
	//3、打开一个新的日志文件
	fileObj,err := os.OpenFile(logName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n",err)
		return nil,err
	}
	//4、将新打开的日志文件对象赋值给 f.fileObj
	return fileObj,nil
}

//记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		//切割
		if f.checkFile(f.fileObj) {
			newFile,err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s-%s-%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			if f.checkFile(f.errFileObj) {
				newFile,err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			//如果要记录的日志大于等于ERROR级别，还要往err日志中再记录一次
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s-%s-%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
