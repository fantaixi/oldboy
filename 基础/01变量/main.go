package main

import "fmt"

var (
	name string
	age int
	isOK bool
)
func main() {
	name = "a"
	age = 18
	isOK = false
	fmt.Print(name+"\n") //在终端输出要打印的类容
	fmt.Printf("age:%d \n",age)  // 格式化输出
	fmt.Println(isOK)  //换行输出
}
