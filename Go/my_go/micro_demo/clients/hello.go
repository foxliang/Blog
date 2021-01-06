package main

import (
	"context"
	"fmt"
	"fox/micro/proto"
	"github.com/micro/go-micro/v2"
	"time"
)

func main() {
	start := time.Now()

	service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
	service.Init()
	helloService := proto.NewHelloService("hello", service.Client())
	res, err := helloService.Ping(context.TODO(), &proto.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)

	end := time.Since(start)
	fmt.Println("Since", end)
}
