package main

import (
	"context"
	"fmt"
	"fox/micro/proto"
	"github.com/micro/go-micro/v2"
)

type Hello struct{}

func (h *Hello) Ping(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello"), // 服务名称
	)
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(Hello))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
