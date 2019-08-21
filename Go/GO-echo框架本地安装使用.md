# 1.Echo简介
Echo是个快速的HTTP路由器（零动态内存分配），也是Go的微型Web框架。其具备快速HTTP路由器、支持扩展中间件，同时还支持静态文件服务、WebSocket以及支持制定绑定函数、制定相应渲染函数，并允许使用任意的HTML模版引擎。

网址：https://echo.labstack.com/

GitHub：https://github.com/labstack/echo

中文网址：http://go-echo.org/

框架比较：
![image](https://img-blog.csdnimg.cn/20190821112109664.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)

# 2.安装

```
$ cd <PROJECT IN $GOPATH>
$ go get -u github.com/labstack/echo/...
```
但是这样是安装不成功的，

![image](https://img-blog.csdnimg.cn/2019082110161321.png)


这里我们缺少crypto组件，需要下载 https://github.com/golang/crypto
下载完成解压到GO的目录下 我这里是：D:\Go\src\golang.org\x  然后再试下

![image](https://img-blog.csdnimg.cn/20190821101753118.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)


这次又提示缺失net组件，不慌我们还去go的github上找：https://github.com/golang/net
下载完成解压到GO的目录下 我这里是：D:\Go\src\golang.org\x  

![image](https://img-blog.csdnimg.cn/20190821131117893.png)


再次下载

![image](https://img-blog.csdnimg.cn/20190821131105489.png0)


 

好了，开启你们的echo之旅吧，顺带说一句，以后再碰到无法安装 golang.org/x/***/ 之类的错误，只需要到 https://github.com/golang/*** 去下载，放到${GOROOT}/src/golang/x/**目录下就OK了
