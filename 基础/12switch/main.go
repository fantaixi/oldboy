package main

import "fmt"

func main() {
	//finger := 3
	//switch finger {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("中指")
	//case 4:
	//	fmt.Println("无名指")
	//case 5:
	//	fmt.Println("小拇指")
	//default:
	//	fmt.Println("无效的输入！")
	//}

	//变种1
	//此时n的作用域只在switch
	//switch n := 7; n {
	//case 1, 3, 5, 7, 9:
	//	fmt.Println("奇数")
	//case 2, 4, 6, 8:
	//	fmt.Println("偶数")
	//default:
	//	fmt.Println(n)
	//}

	//switch中使用判断
	n := 7
	switch {
	case n%2 == 0:
		fmt.Println("偶数")
	case n%2 != 0:
		fmt.Println("奇数")
	default:
		fmt.Println("你有病")
	}

	//穿透  fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	//输出 a  b
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
