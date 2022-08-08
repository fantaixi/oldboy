package main

import "fmt"

//常量  定义之后不可以修改
const pi = 3.1415926

//批量声明
const (
	statusOk = 200
	notFound = 404
)
//批量声明如果后面没有写值，那么默认和上一个值一样
const (
	n1 = 100
	n2
	n3
)

//iota
//在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
//使用iota能简化定义，在定义枚举时很有用。
const (
	a1 = iota  //0
	a2         //1
	a3         //2
)

const (
	b1 = iota  //0
	b2         //1
	_
	b3         //3
)

//插队
//const中每新增一行常量声明将使iota计数一次
const (
	c1 = iota  //0
	c2 = 100
	c3 = iota  //2
	c4         //3
)

//多个常量声明在一行
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
func main() {
	fmt.Println(n3)
	fmt.Println(a2)
	fmt.Println(b3)
	fmt.Println(c3,c4)
	fmt.Println(a,b,c,d,e,f)
}
