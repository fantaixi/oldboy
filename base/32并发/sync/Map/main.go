package main
/*
Go语言的sync包中提供了一个开箱即用的并发安全版 map——sync.Map。
开箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。
同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

方法名																						功能
func (m *Map) Store(key, value interface{})												存储key-value数据
func (m *Map) Load(key interface{}) (value interface{}, ok bool)						查询key对应的value
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)		查询或存储key对应的value
func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)			查询并删除key
func (m *Map) Delete(key interface{})													删除key
func (m *Map) Range(f func(key, value interface{}) bool)								对map中的每个key-value依次调用f
 */
import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全的map
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 存储key-value
			value, _ := m.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}