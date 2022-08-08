package main

import "fmt"

type person struct {
	name    string
	age     int
	hobbies []string
}

func main() {
	var p person
	p.name = "小林"
	p.age = 18
	p.hobbies = []string{"s", "asda", "sad"}
	fmt.Println(p)

	//匿名结构体:多用于临时场景
	var s struct{
		name string
		age int
	}
	s.name = "aaa"
	s.age = 18
	fmt.Printf("%T  %v\n",s,s)
}
