package main

import (
	"flag"
	"tiktok/config"
	"tiktok/db"
)

var FlagConfig = flag.String("f", "./config/config.yaml", "choose config file (.yaml)")

func main() {
	serverConfig := config.ServerConfig{}

	ctx := config.GetConfigMessageFromYaml(FlagConfig, &serverConfig)
	db.InitDb(ctx)
}
