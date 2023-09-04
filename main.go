package main

import (
	"flag"
	"fmt"
	"tiktok/config"
)

var FlagConfig = flag.String("f", "./config/config.yaml", "choose config file (.yaml)")

func main() {
	serverConfig := config.ServerConfig{}
	config.GetConfigMessageFromYaml(FlagConfig, &serverConfig)
	fmt.Println(serverConfig)
}
