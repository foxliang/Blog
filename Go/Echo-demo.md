[Echo框架](https://github.com/weilyf2017/Blog/blob/master/Go/Echo%E6%A1%86%E6%9E%B6%E6%9C%AC%E5%9C%B0%E5%AE%89%E8%A3%85%E4%BD%BF%E7%94%A8.md)

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


# 2.定义get请求，url模式为：/users/:id

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
