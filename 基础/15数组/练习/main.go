package main

import "fmt"

/*
1、求数组[1, 3, 5, 7, 8]所有元素的和
2、找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
 */
func main() {
	num := []int{1, 3, 5, 7, 8}
	res := SumToal(num,8)
	fmt.Println(res)
}
func Sum(num []int) int {
	sum:=0
	for i := 0; i < len(num); i++ {
		sum += i
	}
	return sum
}
func SumToal(num []int, target int) []int {
	res := []int{}
	for i := 0; i < len(num); i++ {
		for j := i+1; j < len(num); j++ {
			if num[i] + num[j] == target {
				res = append(res,i,j)
			}
		}
	}
	return res
}

//用map去重
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x];ok {
			return []int{p,i}
		}
		hashTable[x] = i
	}
	return nil
}