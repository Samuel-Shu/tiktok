package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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
