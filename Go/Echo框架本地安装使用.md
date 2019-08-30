# 1.Echo简介
Echo是个快速的HTTP路由器（零动态内存分配），也是Go的微型Web框架。其具备快速HTTP路由器、支持扩展中间件，同时还支持静态文件服务、WebSocket以及支持制定绑定函数、制定相应渲染函数，并允许使用任意的HTML模版引擎。

网址：https://echo.labstack.com/

GitHub：https://github.com/labstack/echo

中文网址：http://go-echo.org/

框架比较：
![image](https://img-blog.csdnimg.cn/20190821112109664.png)

# 2.安装

```
$ cd <PROJECT IN $GOPATH>
$ go get -u github.com/labstack/echo/...
```
但是这样是安装不成功的，

![image](https://img-blog.csdnimg.cn/2019082110161321.png)


这里我们缺少crypto组件，需要下载 https://github.com/golang/crypto
下载完成解压到GO的目录下 我这里是：D:\Go\src\golang.org\x  然后再试下

![image](https://img-blog.csdnimg.cn/20190821101753118.png)


这次又提示缺失net组件，不慌我们还去go的github上找：https://github.com/golang/net
下载完成解压到GO的目录下 我这里是：D:\Go\src\golang.org\x  

![image](https://img-blog.csdnimg.cn/20190821131117893.png)


再次下载

![image](https://img-blog.csdnimg.cn/20190821131105489.png)


好了，开启你们的echo之旅吧，顺带说一句，以后再碰到无法安装 golang.org/x/***/ 之类的错误，只需要到 https://github.com/golang/*** 去下载，放到${GOROOT}/src/golang/x/**目录下就OK了

# 3.路由&控制器
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
