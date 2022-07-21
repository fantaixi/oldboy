package main

import "fmt"

func main() {
	s1 := []string{"a","b","c"}
	//s1[3] = "d" //会越界
	fmt.Printf("s1:%v len:%d cap:%d\n",s1,len(s1),cap(s1))

	//调用append函数必须使用原来的切片变量接收返回值
	//append追加元素，如果需要扩容的话，会新开辟一段内存空间，此时就会把原来的数据复制过去，因此必须用原来的变量名来接收
	s1 = append(s1,"d")
	fmt.Printf("s1:%v len:%d cap:%d\n",s1,len(s1),cap(s1))

	//扩容规则
	/*
	首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
	即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
	 */
	//相关源码如下  runtime/slice.go
	/*
	newcap := old.cap
		doublecap := newcap + newcap
		if cap > doublecap {
			newcap = cap
		} else {
			if old.cap < 1024 {
				newcap = doublecap
			} else {
				// Check 0 < newcap to detect overflow
				// and prevent an infinite loop.
				for 0 < newcap && newcap < cap {
					newcap += newcap / 4
				}
				// Set newcap to the requested cap when
				// the newcap calculation overflowed.
				if newcap <= 0 {
					newcap = cap
				}
			}
		}
	 */
	/*
	小于1024容量就变成原来的两倍，大于的话就每次增加四分之一
	 */

	ss := []string{"a","w","q"}
	s1 = append(s1,ss...)  // ... 表示拆开
	fmt.Printf("s1:%v len:%d cap:%d\n",s1,len(s1),cap(s1))
}
