package main

import (
	"context"
	"fmt"
	"time"
)

/*
通常健壮的程序都是要设置超时时间的，避免因为服务端长时间响应消耗资源，
所以一些web框架或rpc框架都会采用withTimeout或者withDeadline来做超时控制，
当一次请求到达我们设置的超时时间，就会及时取消，不在往下执行。
withTimeout和withDeadline作用是一样的，就是传递的时间参数不同而已，
他们都会通过传入的时间来自动取消Context，这里要注意的是他们都会返回一个cancelFunc方法，
通过调用这个方法可以达到提前进行取消，不过在使用的过程还是建议在自动取消后也调用cancelFunc去停止定时减少不必要的资源浪费。

withTimeout、WithDeadline不同在于WithTimeout将持续时间作为参数输入而不是时间对象，
这两个方法使用哪个都是一样的，看业务场景和个人习惯了，因为本质withTimout内部也是调用的WithDeadline。
 */

func main() {
	//HttpHandler()  //达到超时时间终止接下来的执行
	HttpHandler1()  //没有达到超时时间终止接下来的执行
}
func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(),3*time.Second)
}
func HttpHandler() {
	ctx,cancle := NewContextWithTimeout()
	defer cancle()
	deal(ctx)
}
func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1*time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n",i)
		}
	}
}

// -------------------------------------------------------------------------------
func NewContextWithTimeout1() (context.Context,context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3 * time.Second)
}

func HttpHandler1()  {
	ctx, cancel := NewContextWithTimeout1()
	defer cancel()
	deal1(ctx, cancel)
}

func deal1(ctx context.Context, cancel context.CancelFunc)  {
	for i:=0; i< 10; i++ {
		time.Sleep(1*time.Second)
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			cancel()
		}
	}
}