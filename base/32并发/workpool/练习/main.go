package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
1、开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
2、开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3、主 goroutine 从resultChan取出结果并打印到终端输出
 */

type Job struct {
	value int64
}

type Result struct {
	job *Job
	sum int64
}

var wg sync.WaitGroup
var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func producer(job chan<- *Job) {
	defer wg.Done()
	for {
		rand.Seed(time.Now().UnixNano())
		newJob := &Job{
			value: rand.Int63(),
		}
		job <- newJob
		time.Sleep(time.Microsecond * 500)
	}
}

func consumer(z1 <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	for {
		job1 := <-z1
		result := int64(0)
		n := job1.value
		for n > 0 {
			result += n % 10
			n = n / 10
		}
		newResult := &Result{
			job: job1,
			sum: result,
		}
		resultChan <- newResult
	}
}
func main() {
	wg.Add(1)
	go producer(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go consumer(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n",result.job.value,result.sum)
	}
	wg.Wait()
}