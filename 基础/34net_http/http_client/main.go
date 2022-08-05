package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,err := http.Get("http://127.0.0.1:9090/xxx?name=sb&age=18")
	if err != nil {
		fmt.Printf("get url failed,err:%v",err)
		return
	}
	//从resp中把服务端返回的数据读取出来
	b,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get body failed,err:%v",err)
		return
	}
	fmt.Println(string(b))
}
