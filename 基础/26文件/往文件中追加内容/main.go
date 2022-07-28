package main

import (
	"fmt"
	"io"
	"os"
)

/*
http://c.biancheng.net/view/5729.html
https://blog.csdn.net/qq_44733706/article/details/113143691
 */
func f1() {
	//打开要操作的文件
	fileObj,err := os.OpenFile("./sb.txt", os.O_RDWR|os.O_CREATE,0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v",err)
		return
	}
	//defer fileObj.Close() //将文件关闭
	// 因为没有办法在文件中间插入内容，需要创建临时文件
	tmpFile, err := os.OpenFile("./sb.tmp",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("creat file failed: ", err)
		return
	}
	//读取原文件写入临时文件
	var ret [1]byte
	n,err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read file failed,err:%v",err)
		return
	}
	// 写入临时文件
	fmt.Println(n)
	fmt.Println(string(ret[:]))
	fmt.Println("1",string(ret[:n]))
	tmpFile.Write(ret[:n])
	// 再写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	//紧接着把源文件后续内容写入临时文件
	var x [1024]byte
	//	for{
	h,err := fileObj.Read(x[:])
	fmt.Println("22222",string(x[:]))
	fmt.Println("333",h)
	fmt.Println("444",string(x[:h]))
	if err == io.EOF {
		fmt.Println("end pf file",x[:h])
		tmpFile.Write(x[:h])
		return
		//break
	}
	if err != nil {
		fmt.Printf("read from file failed: %v",err)
		return
	}
	fmt.Println("写入文件中：",string(x[:h]))
	tmpFile.Write(x[:h])
	//	}
	//源文件后续写入临时文件中
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp","./sb.txt")
}
func main() {
	f1()
}
