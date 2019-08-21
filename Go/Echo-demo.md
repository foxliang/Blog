[Echo框架](https://github.com/weilyf2017/Blog/blob/master/Go/GO-echo%E6%A1%86%E6%9E%B6%E6%9C%AC%E5%9C%B0%E5%AE%89%E8%A3%85%E4%BD%BF%E7%94%A8.md)

# 1.编写 Hello, World!

创建 server.go 文件

```
package main

import (
	"net/http"
    
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```
启动服务

```
$ go run server.go

```

![image](https://img-blog.csdnimg.cn/20190821164530844.png)

在浏览器上输入：http://127.0.0.1:1323/

![image](https://img-blog.csdnimg.cn/20190821164619218.png)


# 2.路由
echo框架的路由定义如下：

```
//定义post请求, url为：/users, 绑定saveUser控制器函数
e.POST("/users", saveUser)

//定义get请求，url模式为：/users/:id  （:id是参数，例如: /users/10, 会匹配这个url模式），绑定getUser控制器函数
e.GET("/users/:id", getUser)

//定义put请求
e.PUT("/users/:id", updateUser)

//定义delete请求
e.DELETE("/users/:id", deleteUser)
```
