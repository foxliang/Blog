package main

import (
	"go.uber.org/zap/zapcore"
	"go_test_log/log"
	"go_test_log/route"
)

func main() {

	log.Init("/runtime/log/go-test-log.txt", &log.LogLevel{Level: zapcore.DebugLevel})
	// start api server
	route.NewServer().Start()
}
