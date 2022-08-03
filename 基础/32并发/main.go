package main

import (
	"fmt"
	"time"
)

func f1(i int) {
	fmt.Println("aaa",i)
}
//程序启动后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 1000; i++ {
		//go f1(i)  //开启一个单独的goroutine去执行f1()
		go func(i int) {
			fmt.Println(i)  //用的是函数参数的i，不是外面的i
		}(i)
	}
	fmt.Println("....")
	time.Sleep(1)
	//main函数结束了，由main函数启动的goroutine也结束了
}
