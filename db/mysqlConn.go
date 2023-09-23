package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"tiktok/config"
)

func InitMysql(c context.Context) *gorm.DB {
	conf := c.Value("config").(*config.ServerConfig)
	dsn := conf.Name + ":" + conf.Password + "@tcp(" + conf.Ip + ":" + conf.Port + ")/" + conf.DatabaseName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("mysql connect failed!", err)
	}
	return db
}
