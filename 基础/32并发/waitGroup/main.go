package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	/*
	Seed使用提供的种子值将默认的Source初始化为一个确定的状态。
	如果不调用Seed，生成器的行为就像用Seed(1)做种子一样。
	在除以2³¹-1时有相同余数的种子值会产生相同的伪随机序列。
	与Rand.Seed方法不同，Seed对于并发使用是安全的。
	 */
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()  //伪随机数
		r2 := rand.Intn(10)  // [0,10)
		fmt.Println(r1,r2)
	}
}
var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}
func main() {
	//f()
	//wg.Add(10)  //直接让计数器加10
	for i := 0; i < 10; i++ {
		wg.Add(1) //执行一次计数器加1
		go f1(i)
	}
	//如何知道这10个goroutine都结束了？
	wg.Wait()  //等待wg的计数器减为0
}
