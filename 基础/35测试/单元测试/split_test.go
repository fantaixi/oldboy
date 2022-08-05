package 单元测试

import (
	"reflect"
	"testing"
)

//测试组
//func Test1Split(t *testing.T) {
//	// 定义一个测试用例类型
//	type test struct {
//		input string
//		sep   string
//		want  []string
//	}
//	// 定义一个存储测试用例的切片
//	tests := []test{
//		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
//		//{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
//	}
//	// 遍历切片，逐一执行测试用例
//	for _, tc := range tests {
//		got := Split(tc.input, tc.sep)
//		if !reflect.DeepEqual(got, tc.want) {
//			t.Errorf("expected:%v, got:%v", tc.want, got)
//		}
//	}
//}

//子测试
//单独跑某一个测试用例   go test -v -run=Test2Split/simple
func Test2Split(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

//测试覆盖率
// go test -cover

//Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。
// go test -cover -coverprofile=c.out
// go tool cover -html=c.out   使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。

//基准测试
