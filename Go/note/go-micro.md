# 一、go-micro是什么

go-micro是基于Go语言实现的插件化RPC微服务框架，与go-kit，kite等微服务框架相比，它具有易上手、部署简单、工具插件化等优点。

go-micro框架提供了服务发现、负载均衡、同步传输、异步通信以及事件驱动等机制，它尝试去简化分布式系统间的通信，让我们可以专注于自身业务逻辑的开发。所以对于新手而言，go-micro是个不错的微服务实践的开始。

# 二、go-micro 架构

### 2.1 分层架构


go-micro是组件化的框架，每一个基础功能都是一个interface，方便扩展。同时，组件又是分层的，上层基于下层功能向上提供服务，整体构成go-micro框架。

go-micro的组件包括：

- Registry组件：服务发现组件，提供服务发现机制：解析服务名字至服务地址。目前支持的注册中心有consul、etcd、 zookeeper、dns、gossip等

- Selector组件：构建在Registry之上的客户端智能负载均衡组件，用于Client组件对Registry返回的服务进行智能选择。

- Broker组件：发布/订阅组件，服务之间基于消息中间件的异步通信方式，默认使用http方式，线上通常使用消息中间件，如Kafka、RabbitMQ等。

- Transport组件：服务之间同步通信方式。

- Codec组件：服务之间消息的编码/解码。

- Server组件：服务主体，该组件基于上面的Registry/Selector/Transport/Broker组件，对外提供一个统一的服务请求入口。

- Client组件：提供访问微服务的客户端。类似Server组件，它也是通过Registry/Selector/Transport/Broker组件实现查找服务、负载均衡、同步通信、异步消息等功能。

所有以上组件功能共同构成一个go-micro微服务。

### 2.2 微服务之间通信

两个微服务之间的通信是基于C/S模型，即服务发请求方充当Client，服务接收方充当Server。

## 编写一个简单的Hello服务
至此，go-micro框架的编程环境已基本搭建好，接下来就是写代码了。

下面实现一个Hello服务：它接收一个字符串类型参数请求，返回一个字符串问候语：Hello 『参数值』。
### 1）定义API

创建proto/hello.proto文件：

使用protobuf文件来定义服务API接口
```
syntax = "proto3";
service Hello {
    rpc Ping(Request) returns (Response) {}
}
message Request {
    string name = 1;
}
message Response {
    string msg = 1;
}
```
### 2）创建service

创建services/hello.go文件：
```
package main

import (
    "context"
    "fmt"

    proto "winmicro/proto"

    micro "github.com/micro/go-micro"
)

type Hello struct{}

func (h *Hello) Ping(ctx context.Context, req *proto.Request, res *proto.Response) error {
    res.Msg = "Hello " + req.Name
    return nil
}
func main() {
    service := micro.NewService(
        micro.Name("hellooo"), // 服务名称
    )
    service.Init()
    proto.RegisterHelloHandler(service.Server(), new(Hello))
    if err := service.Run(); err != nil {
        fmt.Println(err)
    }
}
```
### 3)模拟client

创建Clients/helloclient.go文件：
```
package main

import (
    "context"
    "fmt"

    proto "winmicro/proto"

    micro "github.com/micro/go-micro"
)

func main() {
    service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
    service.Init()
    helloservice := proto.NewHelloService("hellooo", service.Client())
    res, err := helloservice.Ping(context.TODO(), &proto.Request{Name: "World ^_^"})
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(res.Msg)
}
```
### 运行Hello服务
启动consul之后

执行micro list service 查看当前已有服务：


> micro list service

consul

执行go run services/hello.go命令，启动hellooo服务：
```
>go run services/hello.go
2021-01-06 14:04:53  file=v2@v2.9.1/service.go:200 level=info Starting [service] hello
2021-01-06 14:04:53  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:42561
2021-01-06 14:04:53  file=grpc/grpc.go:697 level=info Registry [mdns] Registering node: hello-62b13e3b-7b6a-4658-82d4-1504fb9815a8
```
再次执行micro list service 查看当前已有服务：

> micro list services

consul

hellooo

即hellooo服务已启动

执行go run clients/hello.go命令

```
Hello World ^_^
```
