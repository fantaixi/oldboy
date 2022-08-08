package main

import (
	"encoding/json"
	"fmt"
)

/*
序列化: 结构体  --->  JSON
反序列化: JSON  ---> 结构体
 */
type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func main() {
	p1 := person{
		Name: "大傻",
		Age: 18,
	}
	//序列化
	b,err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n",string(b))
	fmt.Println(string(b))
	//反序列化
	str := `{"name":"大傻...","age":18}`
	var p2 person
	json.Unmarshal([]byte(str),&p2) //传指针是为了能在json.Unmarshal中修改p2的值
	fmt.Printf("%#v\n",p2)
}
