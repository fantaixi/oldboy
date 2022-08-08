package asd

import "fmt"

func init() {
	fmt.Println("外面的init函数执行了...")
}
func Add(x, y int) int {
	return x+y
}
