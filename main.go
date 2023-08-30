package main

import (
	"TikTok/db"
	"TikTok/router"
)

func main() {
	//连接数据库，连接redis
	err := db.Init()
	if err != nil {
		panic("数据库连接错误")
	}

	defer db.MysqlClose()

	//db.InitRedis()

	//注册路由
	r := router.InitRouter()

	err = r.Run(":8080")
	if err != nil {
		panic("服务器启动失败")
	}
}
