package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

//需要收集的日志的配置信息
type LogEntry struct {
	Path string `json:"path"`  //日志存放的
	Topic string `json:"topic"` //日志要发往kafka中的哪个topic
}
//初始化ETCD的函数
func Init(addr string,timeout time.Duration) (err error){
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

//从etcd中根据key获取配置项
func GetConf(key string) (logEntryConf *[]LogEntry,err error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value,&logEntryConf)
		if err != nil {
			fmt.Printf("Unmarshal etcd value failed,err:%v\n",err)
			return
		}
	}
	return
}
