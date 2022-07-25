package main

import "fmt"

//自定义类型和类型别名

//type后面跟的是类型
type myInt int  //自定义类型
type yourInt = int  //类型别名

func main() {
	var n myInt
	fmt.Printf("%T\n",n)

	var m yourInt
	fmt.Printf("%T\n",m)

	//比如rune是int32的别名
	var c rune
	fmt.Printf("%T\n",c)
}
