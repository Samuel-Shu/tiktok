package db

import "github.com/redis/go-redis/v9"

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
}

func RedisClose() {
	err := RDB.Close()
	if err != nil {
		panic("redis连接断开失败")
	}
}
