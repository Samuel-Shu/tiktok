package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"tiktok/config"
)

func InitRedis(ctx context.Context) *redis.Client {
	conf := ctx.Value("config").(*config.ServerConfig)
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.RdbPassword, // no password set
		DB:       conf.Num,         // use default DB
	})
	return rdb
}
