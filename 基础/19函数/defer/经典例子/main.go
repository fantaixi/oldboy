package main

import "fmt"

/*
在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
 */
func f1() int {
	x := 5
	defer func() {
		x++  //返回的是匿名的变量，修改的是x不是返回值，所以结果是5
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++  // 返回的是x，x被修改，所以结果是6
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++  //同 f1
	}()
	return x  // 返回值 = y = x = 5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)  //这里将x当做参数传进去，改变的是副本 没有改变原来x的值，所以x还是5
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
