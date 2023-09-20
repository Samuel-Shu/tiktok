package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DbStruct struct {
	Mysql *gorm.DB
	Redis *redis.Client
}

var Db = &DbStruct{}

func InitDb(ctx context.Context) {
	Db = &DbStruct{Mysql: InitMysql(ctx), Redis: InitRedis(ctx)}
}
