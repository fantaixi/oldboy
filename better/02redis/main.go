package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

func initRedis() (err error) {
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println(pong)
	fmt.Println("成功连接redis")
	return
}
func main() {
	err := initRedis()
	if err != nil {
		panic(err)
	}

	// zsetDemo 操作zset示例
	// key
	zsetKey := "rank"
	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err = rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
}