package main

import "fmt"

/*
内置函数					介绍
close				主要用来关闭channel
len					用来求长度，比如string、array、slice、map、channel
new					用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
make				用来分配内存，主要用来分配引用类型，比如chan、map、slice
append				用来追加元素到数组、slice中
panic和recover		用来做错误处理

Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
panic可以在任何地方引发，但recover只有在defer调用的函数中有效。

注意：
1、recover()必须搭配defer使用。
2、defer一定要在可能引发panic的语句之前定义。
*/
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		fmt.Println(err)
		if err != nil {
			fmt.Println("释放数据库连接。。。")
		}
	}()
	panic("panic in B")  //程序崩溃退出
	fmt.Println("func B")  // Unreachable code  无法到达的代码
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
