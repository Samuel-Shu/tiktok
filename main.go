package main

import (
	"flag"
	"tiktok/config"
	"tiktok/db"
	"tiktok/routers"
)

var FlagConfig = flag.String("f", "./config/config.yaml", "choose config file (.yaml)")

func main() {
	serverConfig := config.ServerConfig{}

	ctx := config.GetConfigMessageFromYaml(FlagConfig, &serverConfig)
	db.InitDb(ctx)
	routers.InitRouter()
}
