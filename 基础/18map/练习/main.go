package main

import (
	"fmt"
	"strings"
)

/*
写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
 */
func main() {
	s := "how do you do"
	TotalStr(s)
}
func TotalStr(s string) {
	res := make(map[string]int)
	s2 := strings.Split(s," ")
	for _, v := range s2 {
		res[v] ++
	}
	fmt.Println(res)
}

/*
观察下面代码，写出最终的打印结果。
 */
/*
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)  // [1 2 3]
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)  //[1 3]
	fmt.Printf("%+v\n", m["q1mi"])//  这里会把前面的2覆盖掉 [1 3 3]
}
 */