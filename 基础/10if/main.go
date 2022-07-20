package main

import "fmt"

func main() {
	//此时score的作用域只在if里面
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
