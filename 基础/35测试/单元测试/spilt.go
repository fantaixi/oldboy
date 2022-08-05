package 单元测试
/*
在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
类型					格式					作用
测试函数		函数名前缀为Test			测试程序的一些逻辑行为是否正确
基准函数		函数名前缀为Benchmark		测试函数的性能
示例函数		函数名前缀为Example		为文档提供示例文档
 */
import "strings"

// ababdf,b  -> aadf
func Split(str string, sep string) []string {
	var res []string
	index := strings.Index(str,sep)
	for index >= 0 {
		res = append(res,str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str,sep)
	}
	res = append(res,str)
	return res
}
