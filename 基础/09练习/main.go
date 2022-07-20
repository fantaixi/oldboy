package main

import "fmt"

//编写代码统计出字符串"hello沙河小王子"中汉字的数量。
func main() {
	s := "hello沙河小王子"
	num := TotalNum(s)
	fmt.Println(num)

	CFB()
}

func TotalNum(s string) int {
	sum := 0
	for _, val := range s {
		if val > 'z' {
			sum++
		}
	}
	return sum
}

func CFB() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		fmt.Println()
	}
}
