package main

import "fmt"

/*
interface : 关键字
interface{}: 空接口类型
 */
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{},16)
	m1["name"] = "aaa"
	m1["age"] = 18
	m1["isBoy"] = true
	m1["hobbies"] = [...]string{"唱","跳","rap"}
	fmt.Println(m1)
	show(m1)
	show(nil)
}
