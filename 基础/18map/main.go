package main

import "fmt"

/*
map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
定义语法如下：
map[KeyType]ValueType
*/
func main() {
	var m1 map[string]int
	//先初始化
	m1 = make(map[string]int, 10)
	m1["a"] = 18
	m1["b"] = 19
	fmt.Println(m1)

	//value,ok := m1["c"]
	//if !ok {
	//	fmt.Println("没有这个KEY")
	//}else {
	//	fmt.Println(value)
	//}
	//或者简写成
	if value,ok := m1["c"];!ok {
		fmt.Println("没有这个KEY")
	}else {
		fmt.Println(value)
	}

	//删除
	delete(m1,"a")
	fmt.Println(m1)
	//删除不存在的key
	delete(m1,"z")  //不会报错
}
