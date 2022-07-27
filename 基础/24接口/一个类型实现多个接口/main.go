package main

import "fmt"

//一个结构体实现多个接口
//接口还可以嵌套
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move()  {
	fmt.Println("走路...")
}
func (c *cat) eat(food string) {
	fmt.Printf("吃%v...\n",food)
}

func main() {

}
