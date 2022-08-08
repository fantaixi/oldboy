package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*
os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。

func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}

其中：
name：要打开的文件名 flag：打开文件的模式。 模式有以下几种：
模式				含义
os.O_WRONLY	    只写
os.O_CREATE	    创建文件
os.O_RDONLY	    只读
os.O_RDWR	    读写
os.O_TRUNC	    清空
os.O_APPEND	    追加
perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
 */

//打开文件写内容
func openWrite() {
	//0644 在win系统下没用，linux系统有用
	// os.O_WRONLY|os.O_CREATE|os.O_APPEND: 打开文件，只写。没有就创建并追加,并且清空
	fileObj,err := os.OpenFile("./xx.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC,0644)
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n",err)
		return
	}
	//Write
	fileObj.Write([]byte("str\n"))       //写入字节切片数据
	//WriteString
	fileObj.WriteString("hello 小王子\n") //直接写入字符串数据
	fmt.Println("aaaaa")
}

//有缓冲区
func demo2() {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

//直接往文件里面写东西
func demo3() {
	str := "hello 沙河........"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
func main() {
	//openWrite()
	//demo2()
	demo3()
}
