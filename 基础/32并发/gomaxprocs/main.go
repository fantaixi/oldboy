package main

/*
Go 语言的调度器采用的是 GPM 调度模型。
1、G：表示 goroutine，每执行一次go f()就创建一个 G，包含要执行的函数和上下文信息。
2、全局队列（Global Queue）：存放等待运行的 G。
3、P：表示 goroutine 执行所需的资源，最多有 GOMAXPROCS 个。
4、P 的本地队列：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。
新建 G 时，G 优先加入到 P 的本地队列，如果本地队列满了会批量移动部分 G 到全局队列。
5、M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，当 P 的本地队列为空时，M 也会尝试从全局队列或其他 P 的本地队列获取 G。
M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。
Goroutine 调度器和操作系统调度器是通过 M 结合起来的，每个 M 都代表了1个内核线程，操作系统调度器负责把内核线程分配到 CPU 的核上执行。

m:n :把m个goroutine分配给n个操作系统线程去执行。

goroutine初试栈的大小是2kb
 */
import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n",i)
	}
}
func f2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n",i)
	}
}
func main() {
	runtime.GOMAXPROCS(4)  //默认CPU的逻辑核心数，默认跑满整个CPU
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}
