package main

import "fmt"

/*
数组
存放元素的容器
必须指定元素的类型和容量（长度）
长度是类型的一部分
*/
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T   a2:%T\n", a1, a2)

	//初始化方式1
	a1 = [3]bool{false, true}
	fmt.Println(a1)
	//初始化方式2
	a3 := [...]int{1, 2, 5, 47, 8}
	fmt.Println(a3)
	//初始化方式3 根据索引初始化
	a4 := [8]int{0: 1, 7: 9}
	fmt.Println(a4)

	//遍历
	//fori   forr
	//....

	//多维数组
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆

	//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	//要想改变就用指针
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
