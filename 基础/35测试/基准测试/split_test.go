package 单元测试

import "testing"

/*
基准测试
基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：
func BenchmarkName(b *testing.B){
    // ...
}
*/
// go test -bench=Split
//还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

//性能比较函数
// go test -bench=.  运行所有
/*
默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。
最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。
像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。

go test -bench=Fib40 -benchtime=20s
 */
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

//通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。
