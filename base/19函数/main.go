package main
/*
函数是组织好的、可重复使用的、用于执行指定任务的代码块。
Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。
 */
func main() {

}
//参数类型简写：有多个连续的参数类型一致的话，可以省略前面的参数
func f1(x, y int) {

}
//可变长参数
// y 可以传递多个值  y的类型是对应类型的切片  而且必须放在参数的最后，否则识别不了
func f2(x string, y ...int) {

}