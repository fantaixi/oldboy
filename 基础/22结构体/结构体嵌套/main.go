package main

import "fmt"

type address struct {
	provice string
	city string
}
type person struct {
	name string
	age int
	//add address
	address //匿名嵌套结构体
}
type company struct {
	name string
	add address
}
func main() {
	p1 := person{
		name: "aaa",
		age: 18,
		//add
		address: address{
			provice: "四川",
			city: "成都",
		},
	}
	fmt.Println(p1)
	//当两个匿名结构体的字段都一样的时候，必须指定是哪个匿名结构体的字段
	fmt.Println(p1.address.city)
	fmt.Println(p1.city)
}
