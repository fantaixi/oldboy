package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start jobs:%d\n",id,j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end jobs:%d\n",id,j)
		results <- j*2
	}
}
func main() {
	jobs := make(chan int,100)
	result := make(chan int,100)
	//开启三个goroutine
	for i := 1; i <= 3; i++ {
		go worker(i,jobs,result)
	}
	//5个任务
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	//输出结果
	for i := 1; i <= 5; i++ {
		<- result
	}
}
