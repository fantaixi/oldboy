package main

import (
	"context"
	"fmt"
	"time"

)

/*
日常业务开发中我们往往为了完成一个复杂的需求会开多个gouroutine去做一些事情，
这就导致我们会在一次请求中开了多个goroutine确无法控制他们，
这时我们就可以使用withCancel来衍生一个context传递到不同的goroutine中，
当我想让这些goroutine停止运行，就可以调用cancel来进行取消。
 */

func main()  {
	ctx,cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(10*time.Second)
	cancel() // 当我们取完需要的整数后调用cancel
	time.Sleep(1*time.Second)
}

func Speak(ctx context.Context)  {
	for range time.Tick(time.Second){
		select {
		case <- ctx.Done():
			fmt.Println("我要闭嘴了")
			return  // return结束该goroutine，防止泄露
		default:
			fmt.Println("balabalabalabala")
		}
	}
}
