package main

import "fmt"

func main() {
	a1 := []int{1,2,3}
	a2 := a1
	a3 := make([]int,len(a1))
	copy(a3,a1)
	fmt.Println(a1,a2,a3)
	a1[0] = 100
	fmt.Println(a1,a2,a3)  //此时a3不受影响

	//Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。
	//删除a1中的2
	a1 = append(a1[:1],a1[2:]...)
	fmt.Println(a1)
	fmt.Println("--------------------------")

	//1、切片不保存具体的值
	//2、切片对应一个底层数组
	//3、底层数组都是占用一块连续的内存
	x1 := [...]int{1,3,5}
	s1 := x1[:]
	fmt.Println(s1,len(s1),cap(s1))
	//删除切片中的3
	s1 = append(s1[:1],s1[2:]...)  //这里修改了底层的数组
	fmt.Println(s1,len(s1),cap(s1))
	//此时原数组中的变化？
	//因为切片本质就是一个指针，指针将原数组中的3变成了5，原来的5没有改变，所以结果就是 [1 5 5]
	fmt.Println(x1)  //[1 5 5]
}
