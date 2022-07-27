package main

import "fmt"

/*
想要从接口值中获取到对应的实际值需要使用类型断言，其语法格式如下。
x.(T)
其中：
x：表示接口类型的变量
T：表示断言x可能是的类型。
 */

//方式1：
// justifyType 对传入的空接口类型变量x进行类型断言
func typeOf(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string,value is %v\n", v)
	case int:
		fmt.Printf("x is a int,value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool,value is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
//方式2:
func typeOf1(a interface{}) {
	fmt.Printf("%T\n",a)
	str,ok := a.(string)
	if !ok{
		fmt.Println("不是字符串")
	}else {
		fmt.Println("是一个字符串",str)
	}
}
func main() {
	typeOf("a")
	typeOf(122)
	typeOf(int64(800))
	typeOf(false)

	typeOf1("aasda")
}
