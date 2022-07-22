package main

import "fmt"

//defer  类似于栈  先进后出
/*
在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
 */
func deferDemo() {
	fmt.Println(1)
	defer fmt.Println(2)  //延迟到函数即将返回的时候再执行
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println(5)
}
func main() {
	deferDemo()  // 1 5 4 3 2
}
