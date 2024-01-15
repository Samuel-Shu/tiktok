package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mini-tiktok/service/core/internal/config"
)

func InitMysql(dataSource string) *gorm.DB {
	engine, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为 Info
	})
	if err != nil {
		log.Printf("failed to connect database:%v", err)
	}
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password, // no password set
		DB:       c.Redis.DB,       // use default DB
	})
	return rdb
}
