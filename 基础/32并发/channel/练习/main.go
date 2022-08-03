package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
1、启动一个goroutine，随机生成100个数到ch1
2、启动一个goroutine，从ch1中取值，计算其平方，放到ch2
3、在main中，从ch2取值打印
 */

var wg sync.WaitGroup

func f1(ch1 chan int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		ch1 <- rand.Intn(100)
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for  {
		x, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	//for v := range ch1 {
	//	ch2 <- v*v
	//}
	close(ch2)
}
func main() {
	a := make(chan int,100)
	b := make(chan int,100)
	wg.Add(2)
	go f1(a)
	go f2(a,b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
