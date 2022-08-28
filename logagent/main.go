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
	//2.1从etcd中获取日志收集项的配置信息
	logEntryConf,err := etcd.GetConf(cfg.EtcdConf.Key)
	//2.2用哨兵去监视日志收集项的变化，有变化就通知logagent实现热加载配置
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n",err)
		return
	}
	fmt.Printf("get conf from etcd success,value:%v\n",logEntryConf)
	//for k, v := range logEntryConf {
	//	fmt.Printf("index:%v,value:%v\n",k,v)
	//}
	//3、收集日志项发往kafka
	//3.1循环每一个日志收集项创建TailObj
	//for _, logEntry := range logEntryConf {
	//
	//}
	//3.2发往kafka

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
