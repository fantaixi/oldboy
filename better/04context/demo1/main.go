package main

import (
	"fmt"
	"sync"
	"time"
)
//让子goroutine退出
//方法1
var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done()
	for  {
		fmt.Println("aaa")
		time.Sleep(time.Millisecond*500)
		if notify {
			break
		}
	}
}
func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	//让子goroutine退出
	notify = true
	wg.Wait()
}
