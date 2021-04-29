package main

import (
	"go-mysql/config"
	"go-mysql/route"
)

func main() {
	config.InitCfg()
	// start api server
	route.NewServer().Start()
}
