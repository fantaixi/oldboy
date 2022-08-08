package main

import (
	"fmt"
	"os"
)

func main() {
	// 36flag.exe a b c
	//使用os.Args来获取命令行参数
	fmt.Printf("%#v\n",os.Args) //[]string{"36flag.exe", "a", "b", "c"}
	//获取第一个和第三个参数
	fmt.Println(os.Args[0],os.Args[2]) //36flag.exe b
	fmt.Printf("%T\n",os.Args)  // []string
}
