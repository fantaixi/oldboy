package main

import (
	"time"
)

//测试自己写的日志库
func main() {
	log := NewLog("Debug")
	for  {
		id := 100
		name := "阿萨德"
		log.Debug("这是一条debug日志... id:%d,name:%s",id,name)
		log.Info("这是一条info日志...")
		log.Warning("这是一条warning日志...")
		log.Error("这是一条error日志...")
		log.Fatal("这是一条fatal日志...")
		time.Sleep(time.Second)
	}
}
