package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"oldboy/logagent/conf"
	"oldboy/logagent/etcd"
	"oldboy/logagent/kafka"
	"time"
)

//func run() {
//	//1、读取日志
//	for {
//		select {
//		  case 	line :=  <-taillog.ReadChan():
//			  //2、发送到kafka
//			  kafka.SendToKafka(cfg.KafkaConf.Topic,line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}

var (
	cfg = new(conf.AppConf)
)

//logagent 入口程序
func main() {
	//1、加载配置文件
	//cfg,err := ini.Load("./conf/config.ini")
	err := ini.MapTo(cfg,"./conf/config.ini")
	if err != nil {
		fmt.Printf("init config failed,err:%v\n",err)
		return
	}
	//2、初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v\n",err)
		return
	}
	fmt.Println("init kafka success")
	//2、初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address,time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n",err)
		return
	}
	fmt.Println("init etcd success")
	//3、打开日志文件准备收集
	//err = taillog.Init(cfg.TaillogConf.FileName)
	//if err != nil {
	//	fmt.Printf("init taillog failed,err:%v\n",err)
	//	return
	//}
	//fmt.Println("init tail success")
	////4、具体的业务
	//run()
}
