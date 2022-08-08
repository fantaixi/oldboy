package main

import (
	"fmt"
	"sync"
)

/*
Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。
channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

通道的操作：
1、发送： ch1 <-  1
2、接受： <- ch1
3、关闭： close()
 */

var a chan int
var wg sync.WaitGroup

func noBuffChannel() {
	fmt.Println(a)
	a = make(chan int)  //不带缓冲区，使用前必须初始化，不然就是nil
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <- a
		fmt.Println("后台goroutine从通道中取到",x)
	}()
	a <- 10
	wg.Wait()
}
func buffChannel() {
	fmt.Println(a)
	a = make(chan int,15)  //带缓冲区的初始化
	a <- 10
	fmt.Println(a)
	fmt.Println("10放到通道中了")
	close(a)
}
func main() {
	//noBuffChannel()
	buffChannel()
}
