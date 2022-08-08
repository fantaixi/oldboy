package main

import (
	"fmt"
	"sync"
	"time"
)

//让子goroutine退出
//方法2  用通道
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("aaa")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan: //如果能从通道取到值就退出
			break LOOP
		default:
		}
	}
}
func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	//让子goroutine退出
	exitChan <- true //随便往管道发送一个值
	wg.Wait()
}
