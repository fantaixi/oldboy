package main

import "github.com/go-redis/redis"

var redisDB *redis.Client

//连接池
func initRedis() (err error){
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	//_,err =redisDB.Ping().Result()
	return
}

func main() {

}
