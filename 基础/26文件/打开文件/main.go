package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//没有换缓冲区用read直接读
func ReadFile() {
	fileObj,err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n",err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	//读文件
	//var tem = make([]byte,128)  //指定每次读取的长度
	//或者
	var tem  [128]byte
	for  {
		n ,err := fileObj.Read(tem[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("读取文件失败,err:%v\n",err)
			return
		}
		fmt.Printf("读取了%d个字节\n",n)
		fmt.Println(string(tem[:n]))
		if n < 128 {
			return
		}
	}
}

//有缓冲区用bufio读
func bufioFile() {
	fileObj,err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n",err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for  {
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("读取文件失败,err:%v\n",err)
			return
		}
		fmt.Print(line)
	}
}

//ioutil读取整个文件
func ioutilDemo() {
	ret,err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("读取文件失败,err:%v\n",err)
		return
	}
	fmt.Println(string(ret))
}
func main() {
	//ReadFile()
	//bufioFile()
	ioutilDemo()
}
