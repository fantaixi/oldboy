package main

import "fmt"

/*
引出接口
 */
type cat struct {}
func (c cat) speak() {
	fmt.Println("猫")
}

type dog struct {}
func (d dog) speak() {
	fmt.Println("狗")
}

type speaker interface {
	speak()  //只有实现了speak()方法的变量都是speaker类型
}
func jiao(x speaker) {
	x.speak()
}
func main() {
	var c1 cat
	var d1 dog
	jiao(c1)
	jiao(d1)
}
