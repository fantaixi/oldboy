package main

import "fmt"

func f1(f func()) {
	fmt.Println("This is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x+y)
}

//要求: f1(f2)
func f3(f func(int,int),x, y int) func() {
	temp := func() {
		//fmt.Println(x+y)
		f(x,y)
	}
	return temp
}
func main() {
	ret := f3(f2,100,1)  //这样就把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(ret)
}
