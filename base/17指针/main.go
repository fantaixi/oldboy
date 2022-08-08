package main

import "fmt"

/*
Go语言中不存在指针操作，只需要记住两个符号：
1、&: 取地址
2、*: 根据地址取值
 */

/*
new与make的区别
1、二者都是用来做内存分配的。
2、make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3、而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。（*int、*string....）
 */
func main() {
	//&: 取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n",p)  //*int: int类型的指针

	//*: 根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n",m)  //int

	//new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false

	//make
	//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
	//而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	//因为这三种类型就是引用类型，所以就没有必要返回他们的指针了

}
