package main

import (
	"fmt"
	"sync"
)

/*
互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。
Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。
sync.Mutex提供了两个方法供我们使用。
方法名						功能
func (m *Mutex) Lock()		获取互斥锁
func (m *Mutex) Unlock()	释放互斥锁
 */
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()  //对公共变量加锁
		x += 1
		lock.Unlock()  //不解锁会造成死锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
