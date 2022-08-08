package main

import "fmt"

func main() {
	//基本格式
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//变种1
	//i := 0
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//变种2
	//i := 0
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//死循环
	//for{
	//
	//}

	//for range 循环
	/*
	 通过for range遍历的返回值有以下规律：
		数组、切片、字符串返回索引和值。
		map返回键和值。
		通道（channel）只返回通道内的值。
	 */
	//s := "asadas"
	//for i, v := range s {
	//	fmt.Printf("i=%d v=%c\n",i,v)
	//}
	////只有一个变量的时候打印的是索引
	//for i := range s {
	//	fmt.Printf("i=%d \n",i)
	//}

	//break:直接跳出循环
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//continue:结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
	//还可以配合标签使用，表示开始标签对应的循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue  //不打印5
		}
		fmt.Println(i)
	}
}
