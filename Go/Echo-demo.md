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


# 2.路由&控制器
1.echo框架的路由定义如下：

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
2.控制器

在echo框架中，控制器是一个函数，我们需要根据业务实现各种控制器函数，控制器函数定义如下：

```
//控制器函数只接受一个echo.Context上下文参数
//参数：c 是上下文参数，关联了当前请求和响应，通过c参数我们可以获取请求参数，向客户端响应结果。
func HandlerFunc(c echo.Context) error
例子：

// 路由定义：e.GET("/users/:id", getUser)
// getUser控制器函数实现
func getUser(c echo.Context) error {
  	// 获取url上的path参数，url模式里面定义了参数:id
  	id := c.Param("id")
  	
  	//响应一个字符串，这里直接把id以字符串的形式返回给客户端。
	return c.String(http.StatusOK, id)
}
```
3.定义get请求，url模式为：/users/:id

```
package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/users/:id", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID 来自于url `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
```

执行此方法并访问：http://127.0.0.1:1323/users/fox

输出
```
fox

```
