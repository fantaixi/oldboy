package main

import "fmt"

//结构体是值类型
type person struct {
	name string
	age int
}

func f(p person) {
	p.age = 19
}
func f1(p *person) {
	//(*p).age = 88
	//可以简写成
	p.age = 19
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}
func main() {
	var p  person
	p.name = "aaa"
	p.age = 18
	fmt.Println(p)

	f(p)  //传的是副本，不会改变原来的值
	fmt.Println(p)

	f1(&p)
	fmt.Println(p)

	//结构体类型的内存地址是连在一起的
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}
