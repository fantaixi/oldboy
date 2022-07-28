package main

import (
	"fmt"
	"time"
)

/*
获取当前时间，格式化输出为2017/06/19 20:30:05格式。
编写程序统计一段代码的执行耗时时间，单位精确到微秒。
 */

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow.Format("2006/01/02 15:04:05"))
	startTime := time.Now().UnixNano()
	fmt.Println(startTime)
	CFB()
	endTime := time.Now().UnixNano()
	fmt.Println(endTime)
	fmt.Println(endTime-startTime)
}
func CFB() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		fmt.Println()
	}
}
