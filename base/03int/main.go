package main

import "fmt"

func main() {
	//十进制
	a1 := 101
	fmt.Printf("%d\n",a1)
	fmt.Printf("%b\n",a1)  //转换成二进制
	fmt.Printf("%o\n",a1)  //转换成八进制
	fmt.Printf("%x\n",a1)  //转成十六进制
	//八进制
	a2 := 077
	fmt.Printf("%d\n",a2)
	//十六进制
	a3 := 0x1234567
	fmt.Printf("%d\n",a3)
	//查看变量类型
	fmt.Printf("%T\n",a3)
	//声明int8类型的变量
	a4 := int8(9)
	fmt.Printf("%T\n",a4)
}
