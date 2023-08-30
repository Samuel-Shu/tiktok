package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (err error) {
	dsn := "root:wjp12345@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	return nil
}

func MysqlClose() {
	db, err := DB.DB()
	if err != nil {
		panic("数据库连接关闭失败")
	}
	err = db.Close()
	if err != nil {
		panic("数据库连接关闭失败")
	}

}
