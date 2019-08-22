# Golang 是什么
Go 亦称为 Golang（译注：按照 Rob Pike 说法，语言叫做 Go，Golang 只是官方网站的网址），是由谷歌开发的一个开源的编译型的静态语言。

Golang 的主要关注点是使得高可用性和可扩展性的 Web 应用的开发变得简便容易。（译注：Go 的定位是系统编程语言，只是对 Web 开发支持较好）

## 为何选择 Golang
既然有很多其他编程语言可以做同样的工作，如 Python，Ruby，Nodejs 等，为什么要选择 Golang 作为服务端编程语言？

以下是我使用 Go 语言时发现的一些优点：

1.并发是语言的一部分（译注：并非通过标准库实现），所以编写多线程程序会是一件很容易的事。后续教程将会讨论到，并发是通过 Goroutines 和 channels 机制实现的。

2.Golang 是一种编译型语言。源代码会编译为二进制机器码。而在解释型语言中没有这个过程，如 Nodejs 中的 JavaScript。

3.语言规范十分简洁。所有规范都在一个页面展示，你甚至都可以用它来编写你自己的编译器呢。:smile:

4.Go 编译器支持静态链接。所有 Go 代码都可以静态链接为一个大的二进制文件（译注：相对现在的磁盘空间，其实根本不大），并可以轻松部署到云服务器，而不必担心各种依赖性。


### 第一个 Go 程序
接下来我们来编写第一个 Go 程序 hello.go（Go 语言源文件的扩展是 .go），代码如下：

```
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```

结果：

```
$ go run test.go
Hello, World!
```

官网：https://golang.org / https://golang.google.cn/ 

Github：https://github.com/golang

Go语言中文网：https://studygolang.com/

GO指南：https://tour.go-zh.org/list
