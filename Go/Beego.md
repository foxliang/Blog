# Beego Framework

beego 是一个快速开发 Go 应用的 HTTP 框架，他可以用来快速开发 API、Web 及后端服务等各种应用，是一个 RESTful 的框架，主要设计灵感来源于 tornado、sinatra 和 flask 这三个框架，但是结合了 Go 本身的一些特性（interface、struct 嵌入等）而设计的一个框架。

### 下载安装

beego 包含一些示例应用程序以帮您学习并使用 beego 应用框架。

您需要安装 Go 1.1+ 以确保所有功能的正常使用。

你需要安装或者升级 Beego 和 Bee 的开发工具:

```
$ go get -u github.com/astaxie/beego
$ go get -u github.com/beego/bee
```
为了更加方便的操作，请将 $GOPATH/bin 加入到你的 $PATH 变量中。请确保在此之前您已经添加了 $GOPATH 变量。

#### 如果您还没添加 $GOPATH 变量
```
$ echo 'export GOPATH="$HOME/go"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件
```

#### 如果您已经添加了 $GOPATH 变量
```
$ echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile # 或者 ~/.zshrc, ~/.cshrc, 您所使用的sh对应的配置文件
$ exec $SHELL
```
#### 想要快速建立一个应用来检测安装？

```
$ cd $GOPATH/src
$ bee new hello
$ cd hello
$ bee run
```
#### Windows 平台下输入：

```
>cd %GOPATH%/src
>bee new hello
>cd hello
>bee run
```
这些指令帮助您：

安装 beego 到您的 $GOPATH 中。

在您的计算机上安装 Bee 工具。

创建一个名为 “hello” 的应用程序。

启动热编译。

一旦程序开始运行，您就可以在浏览器中打开 http://localhost:8080/ 进行访问。

![image](https://github.com/foxliang/Blog/blob/master/images/beego.png)



网站：https://beego.me/

Github：https://github.com/beego/bee
