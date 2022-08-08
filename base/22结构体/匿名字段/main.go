package main

import "fmt"

//匿名字段，而且相同类型只能写一个
type person struct {
	string
	int
}
func main() {
	p1 := person{
		"aaa",
		18,
	}
	fmt.Println(p1)
	//没有名字的时候默认把类型当名字
	fmt.Println(p1.string)
}
