package main

import (
	"fmt"
	a "oldboy/base/25包/asd" //可以给包取别名
)

func init() {
	fmt.Println("main包中的函数执行了...")
}
func main() {
	res := a.Add(8,89)
	fmt.Println(res)
}
