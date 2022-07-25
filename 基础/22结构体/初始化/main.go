package main

import "fmt"

//构造函数和初始化
type person struct {
	name string
	age int
}

//当结构体比较大的时候，返回结构体指针能减少开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age: age,
	}
}
func main() {
	p1 := newPerson("aaa",18)
	fmt.Println(p1)
}
