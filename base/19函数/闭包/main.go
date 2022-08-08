package main

import "fmt"

/*
闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用外面的变量
 */

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder(100)
	fmt.Println(f(10)) //110
}

