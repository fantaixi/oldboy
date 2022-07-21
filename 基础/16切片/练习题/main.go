package main

import "fmt"

func main() {
	//长度为5 表示已经有五个元素  再追加就是下面的结果
	var a = make([]int, 5, 10)
	fmt.Println(a)  // [0 0 0 0 0]
	for i := 0; i < 10; i++ {
		a = append(a,  i)
	}
	fmt.Println(a)  // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
}
