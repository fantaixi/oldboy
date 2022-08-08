package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//让子goroutine退出
//方法3  用context
var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
LOOP:
	for {
		fmt.Println("aaa")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //向外暴露一个只读的空结构体  等待上级通知
			break LOOP
		default:
		}
	}
}
func f2(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("qqq")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //向外暴露一个只读的空结构体  等待上级通知
			break LOOP
		default:
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	//让子goroutine退出
	cancel()
	wg.Wait()
}
