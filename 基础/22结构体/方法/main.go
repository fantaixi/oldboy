package main

import "fmt"

/*
方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。
接收者的概念就类似于其他语言中的this或者 self。
方法的定义格式如下：
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
其中，
接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：具体格式与函数定义相同。
 */

type person struct {
	name string
}

func newPerson(name string) *person {
	return &person{
		name: name,
	}
}

//方法
func (p *person) eat() {
	fmt.Printf("%s 要吃饭\n",p.name)
}
func main() {
	p1 := newPerson("aaaa")
	p1.eat()
}
