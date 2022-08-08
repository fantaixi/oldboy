package main

import "fmt"

//可以声明在全局
//var f1 = func(x,y int) {
//	fmt.Println(x+y)
//}
func main() {
	//函数内部无法显示的声明另一个函数
	//要想实现这个功能就用匿名函数
	var f1 = func(x, y int) {
		fmt.Println(x + y)
	}
	f1(3, 9)

	//如果只是调用一次的函数，还可以简写成立即执行函数
	func() {
		fmt.Println("aaa")
	}()
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 2)   //后面的括号表示传给匿名函数的参数
}
