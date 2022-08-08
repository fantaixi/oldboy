package main

import (
	"fmt"
	"strings"
)

// 字符串是用双引号包裹   字符是用单引号
func main() {
	//转义符
	path := "\"E:\\BaiduNetdiskDownload\\后端开发大礼包（来自公众号「后端进阶」）\""
	fmt.Println(path)

	//多行字符串
	s2 := `
		   sad
		fsd
		sadasdn
	`
	fmt.Println(s2)

	//拼接
	name := "aaa"
	age := "18"
	ss1 := fmt.Sprintf("%s%s",name,age)
	fmt.Println(ss1)

	//切割
	ret := strings.Split(path,"\\")
	fmt.Println(ret)

	//是否包含
	fmt.Println(strings.Contains(ss1,"aab"))

	//前缀
	fmt.Println(strings.HasPrefix(ss1,"aaa"))
	//后缀
	fmt.Println(strings.HasSuffix(ss1,"18"))
	//子串出现的位置
	s4 := "abcdbe"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.LastIndex(s4,"b"))
	//join操作:拼接
	fmt.Println(strings.Join(ret,"+"))

	//字符串修改
	a1 := "abc"
	a2 := []rune(a1)  //强制类型转换
	a2[0] = 'b'
	fmt.Println(string(a2))
}
