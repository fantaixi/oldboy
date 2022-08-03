package main

import "fmt"

/*
Select 的使用方式类似于之前学到的 switch 语句，它也有一系列 case 分支和一个默认的分支。
每个 case 分支会对应一个通道的通信（接收或发送）过程。
select 会一直等待，直到其中的某个 case 的通信操作完成时，就会执行该 case 分支对应的语句

Select 语句具有以下特点。
1、可处理一个或多个 channel 的发送/接收操作。
2、如果多个 case 同时满足，select 会随机选择一个执行。
3、对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
 */
func main() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
