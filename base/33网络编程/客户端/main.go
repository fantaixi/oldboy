package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1、与服务端建立连接
	conn,err := net.Dial("tcp","127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed,err:",err)
		return
	}
	//2、通信
	reader := bufio.NewReader(os.Stdin)
	for  {
		fmt.Print("请说话:")
		msg,_ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
