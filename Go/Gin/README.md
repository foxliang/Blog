# Gin 是什么？
Gin是一个用Go语言编写的web框架。它是一个类似于martini但拥有更好性能的API框架, 由于使用了httprouter，速度提高了近40倍。 如果你是性能和高效的追求者, 你会爱上Gin。

## Gin框架介绍
Go世界里最流行的Web框架，[Github](https://github.com/gin-gonic/gin)上有32K+star。 基于httprouter开发的Web框架。 [中文文档](https://gin-gonic.com/zh-cn/docs/)齐全，简单易用的轻量级框架。

#### 下载并安装Gin:
```
go get -u github.com/gin-gonic/gin
```
#### 第一个Gin示例：
```
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
```
将上面的代码保存并编译执行，然后使用浏览器打开127.0.0.1:8080/hello就能看到一串JSON字符串
