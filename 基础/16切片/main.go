package main

import "fmt"

/*
切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
切片是一个引用类型，它的内部结构包含地址、长度和容量。
长度就是元素的个数
切片的容量是指底层数组的切片指向的第一个元素到最后一个元素的数量
 */
func main() {
	//定义一个切片
	var s1 []int
	fmt.Println(s1)
	fmt.Println(s1 == nil)

	//初始化
	s1 = []int{1,2,3}
	fmt.Println(s1)
	fmt.Println(s1 == nil)

	//长度len()和容量cap()
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	//由数组得到切片(左闭右开)
	a1 := []int{1,3,5,7,8,96,4}
	s2 := a1[0:3]
	fmt.Println(s2)
	//还有如下方式
	//s3 := a1[:3]
	//s4 := a1[0:]
	s5 := a1[:]

	//切片再切片
	s6 := s5[:len(s5)-2]
	fmt.Printf("len: %d cap: %d\n",len(s6),cap(s6))
}
