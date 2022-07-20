package main

import "fmt"

func main() {
	//跳出多层循环
	//flag := false
	//for i := 0; i < 10; i++ {
	//	for j := 'A'; j < 'Z'; j++ {
	//		if j == 'C' {
	//			flag = true
	//			break   //跳出内层循环
	//		}
	//		fmt.Printf("%v--%c\n",i,j)
	//	}
	//	if flag {
	//		break  //跳出外层循环
	//	}
	//}

	//使用goto+lable
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx   //跳到指定的标签
			}
			fmt.Printf("%v--%c\n",i,j)
		}
	}
	//lable标签  名字随便起
	xx:
		fmt.Println("over")

	//标签也可以和break和continue一起使用
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}
