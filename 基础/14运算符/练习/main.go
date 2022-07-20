package main

import "fmt"

/*
有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
 */
/*
将所有的数字都异或，那么最终结果就是要找的数字，因为相同的数字异或结果就是0
 */
func main() {
	num := []int{1,5,5,3,3,6,1}
	res := YiHuo(num)
	fmt.Println(res)
}
func YiHuo(num []int) int {
	i:=0
	for _, v := range num {
		i ^= v
	}
	return i
}
