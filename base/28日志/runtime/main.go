package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	/*
	// Caller在调用goroutine的堆栈中报告有关函数调用的文件和行号信息。
	参数skip是要上升的堆栈帧数，其中0表示Caller的调用者。  (由于历史原因，跳过的含义在Caller和Callers之间有所不同）。
	返回值报告的是程序计数器、文件名和相应的调用文件中的行号。
	如果不可能恢复这些信息，则布尔值ok为false。
	 */
	//skip表示在第几层调用，如果在外面有函数调用，那么skip应该为1
	pc,file,line,ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime failed\n")
		return
	}
	//这样才能拿到调用者的名字
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	//拿到文件名字
	fmt.Println(path.Base(file))
	fmt.Println(line)
}
