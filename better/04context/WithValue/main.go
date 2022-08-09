package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)
/*
WithValue携带数据
我们日常在业务开发中都希望能有一个trace_id能串联所有的日志，这就需要我们打印日志时能够获取到这个trace_id，
在python中我们可以用gevent.local来传递，在java中我们可以用ThreadLocal来传递，在Go语言中我们就可以使用Context来传递，
通过使用WithValue来创建一个携带trace_id的context，然后不断透传下去，打印日志时输出即可

基于context.Background创建一个携带trace_id的ctx，然后通过context树一起传递，从中派生的任何context都会获取此值，
我们最后打印日志的时候就可以从ctx中取值输出到日志中。目前一些RPC框架都是支持了Context，所以trace_id的向下传递就更方便了。

在使用withVaule时要注意四个事项：

-不建议使用context值传递关键参数，关键参数应该显示的声明出来，不应该隐式处理，context中最好是携带签名、trace_id这类值。
-因为携带value也是key、value的形式，为了避免context因多个包同时使用context而带来冲突，key建议采用内置类型。
-上面的例子我们获取trace_id是直接从当前ctx获取的，实际我们也可以获取父context中的value，在获取键值对是，
  我们先从当前context中查找，没有找到会在从父context中查找该键对应的值直到在某个父context中返回 nil 或者查找到对应的值。
-context传递的数据中key、value都是interface类型，这种类型编译期无法确定类型，所以不是很安全，所以在类型断言时别忘了保证程序的健壮性。
 */
const (
	KEY = "trace_id"   //全局的id
)

func NewRequestID() string {
	//生成uuid
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	//基于context.Background生成trace_id
	ctx := context.WithValue(context.Background(), KEY,NewRequestID())
	return ctx
}

func PrintLog(ctx context.Context, message string)  {
	//格式化输出
	fmt.Printf("%s|info|trace_id=%s|%s",time.Now().Format("2006-01-02 15:04:05") , GetContextValue(ctx, KEY), message)
}

func GetContextValue(ctx context.Context,k string)  string{
	//类型断言，拿到context的值
	v, ok := ctx.Value(k).(string)
	if !ok{
		return ""
	}
	return v
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "GolangYYDS")
}


func main()  {
	ProcessEnter(NewContextWithTraceID())
}
