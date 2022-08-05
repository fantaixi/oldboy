package main

import (
	"flag"
	"fmt"
)

func main() {
	//使用flag.Type来获取命令行参数
	// flag.exe -name=睿智 -age=18
	//创建一个标志位参数
	//name := flag.String("name","睿智","请输入名字")
	//age := flag.Int("age",1000,"请输入年龄")
	////时间段flag  合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”。
	////使用flag
	//flag.Parse()
	//fmt.Println(*name)
	//fmt.Println(*age)

	//或者使用flag.TypeVar()
	var name string
	var age int
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	fmt.Println(name)
	fmt.Println(age)

	/*
	flag 其他函数
	flag.Args()  //返回命令行参数后的其他参数，以[]string类型
	flag.NArg()  //返回命令行参数后的其他参数个数
	flag.NFlag() //返回使用的命令行参数个数
	 */
}
