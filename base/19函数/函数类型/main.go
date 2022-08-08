package main

import "fmt"

//函数也可以当做一个类型来使用
func f1() {
	fmt.Println("aaa")
}
func f2() int {
	return 10
}
//函数也可以当做参数类型
func f3(x func() int) {
	res := x()
	fmt.Println(res)
}
//函数还可以作为返回值
func f4(x func()int) func(int, int) int {
	res := func(x,y int) int {
		return x+y
	}
	return res
}
func main() {
	a := f1
	fmt.Printf("%T\n",a)  //func()
	//函数类型也是有区别的
	b := f2
	fmt.Printf("%T\n",b)  //func() int

	f3(f2)

	f4(f2)
}
