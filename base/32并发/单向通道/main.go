package main
/*
在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，通常我们会选择在不同的任务函数中对通道的使用进行限制，
比如限制通道在某个函数中只能执行发送或只能执行接收操作。
想象一下，我们现在有Producer和Consumer两个函数，
其中Producer函数会返回一个通道，并且会持续将符合条件的数据发送至该通道，并在发送完成后将该通道关闭。
而Consumer函数的任务是从通道中接收值进行计算，这两个函数之间通过Processer函数返回的通道进行通信。

一般在函数参数进行限制
func X(x chan<- int,y <-chan int){}
 */
import (
"fmt"
)

// Producer2 返回一个接收通道
func Producer2() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer2 参数为接收通道
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch2 := Producer2()
	res2 := Consumer2(ch2)
	fmt.Println(res2) // 25
}
