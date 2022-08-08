package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	//3、与客户端通信
	var temp [128]byte
	for{
		n,err := conn.Read(temp[:])
		if err != nil {
			fmt.Println("read failed,err:",err)
			return
		}
		fmt.Println(string(temp[:n]))
	}
}
func main() {
	//1、本地端口启动服务
	listener,err := net.Listen("tcp","127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp failed,err:",err)
		return
	}
	//2、等待别人连接
	for  {
		conn,err :=  listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:",err)
			return
		}
		go processConn(conn)
	}
}
