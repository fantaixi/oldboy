package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串  --> 数字
	str := "100000"
	//ParseInt 返回的是int64,想要int就讲bitSize设置为0，最后结果强制转换
	ret1,_ := strconv.ParseInt(str,10,64)
	fmt.Printf("%v %T\n",ret1,ret1)

	//或者直接使用Aoti
	ret2,_ := strconv.Atoi(str)
	fmt.Printf("%v %T\n",ret2,ret2)
	// 数字  ---->  字符串
	i := 97
	ret3 := fmt.Sprintf("%d",i)
	fmt.Printf("%#v\n",ret3)
	//或者
	ret4 := strconv.Itoa(i)
	fmt.Printf("%#v\n",ret4)
}
