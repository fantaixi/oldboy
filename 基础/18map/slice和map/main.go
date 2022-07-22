package main

import "fmt"

//map和slice组合
func main() {
	//元素类型为map的切片
	var s1 = make([]map[int]string, 10)
	//s1[0][10] = "a"  // assignment to entry in nil map
	//必选先初始化切片里的map才能继续使用
	s1[0] = make(map[int]string,1)
	s1[0][10] = "a"
	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int,10)
	m1["aaa"] = []int{1,2,3,5}
	fmt.Println(m1)
}
