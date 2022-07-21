package main

import "fmt"

//使用make创建切片
func main() {
	//指定类型、长度和容量
	s1 := make([]int,5,10)
	fmt.Printf("s1=%v len:%d cap:%d\n",s1,len(s1),cap(s1))

	//指定类型、长度
	s2 := make([]int,5)
	fmt.Printf("s2=%v len:%d cap:%d\n",s2,len(s2),cap(s2))

	//切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。
	//切片唯一合法的比较操作是和nil比较。
	//一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	//但是我们不能说一个长度和容量都是0的切片一定是nil，
	/*
	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	 */
	//所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。

	//切片的赋值
	s3 := []int{1,2,3}
	s4 := s3
	fmt.Println(s3,s4)
	s4[0] = -1
	fmt.Println(s3,s4)
}
