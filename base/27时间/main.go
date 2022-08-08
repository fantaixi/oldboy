package main

import (
	"fmt"
	"time"
)
func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//time.Unix()  根据时间戳拿到时间
	ret := time.Unix(1658994522, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	//时间间隔
	fmt.Println(time.Second)

	//now+24小时
	fmt.Println(now.Add(24 * time.Hour))

	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t)  //一秒钟执行一次
	//}

	//格式化  把时间对象转换成字符串
	/*
		time.Format函数能够将一个时间对象格式化输出为指定布局的文本表示形式，
		需要注意的是 Go 语言中时间格式化的布局不是常见的Y-m-d H:M:S，
		而是使用 2006-01-02 15:04:05.000（记忆口诀为2006 1 2 3 4 5）。
	*/
	//2022-07-28
	fmt.Println(now.Format("2006-01-02")) //2022-07-28
	//2022/07/28 15:04:05
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//2022/07/28 15:04:05 PM  十二进制的时间
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	//精确到毫秒
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	//解析字符串格式的时间
	timeObj, err := time.Parse("2006-01-02", "1994-08-25")
	if err != nil {
		fmt.Printf("解析时间失败,err: %v", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Year())

	//算两个时间相减
	nextYear, err := time.Parse("2006-01-02 15:04:05", "1994-08-25 12:01:01")
	if err != nil {
		fmt.Printf("解析时间失败,err: %v", err)
		return
	}
	//nextYear到现在的时间间隔
	d := now.Sub(nextYear)
	fmt.Println(d)

	//sleep
	n := 5
	fmt.Println("开始睡3秒钟")
	//Sleep需要传Duration类型，也就是int64
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("结束了")
}

//时区
func f2() {
	now := time.Now()  //本地的时间
	fmt.Println(now)
	//明天的这个时间
	//按照指定格式去解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2022-07-29 12:01:01")
	//按照东八区的时区和格式解析一个字符串格式的时间
	//根据字符串加载时区
	loc,err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("loca failed,err:%v\n",err)
		return
	}
	//按照指定时区解析时间
	timeObj,err := time.ParseInLocation("2006-01-02 15:04:05", "2022-07-29 12:01:01",loc)
	if err != nil {
		fmt.Printf("ParseInLocation failed,err:%v\n",err)
		return
	}
	fmt.Println(timeObj)
	//时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	//f1()
	f2()
}
