# PHP-FPM 介绍
## CGI 协议与 FastCGI 协议
每种动态语言（ PHP,Python 等）的代码文件需要通过对应的解析器才能被服务器识别，而 CGI 协议就是用来使解释器与服务器可以互相通信。PHP 文件在服务器上的解析需要用到 PHP 解释器，再加上对应的 CGI 协议，从而使服务器可以解析到 PHP 文件。

由于 CGI 的机制是每处理一个请求需要 fork 一个 CGI 进程，请求结束再kill掉这个进程，在实际应用上比较浪费资源，于是就出现了CGI 的改良版本 FastCGI，FastCGI 在请求处理完后，不会 kill 掉进程，而是继续处理多个请求，这样就大大提高了效率。

## PHP-FPM 是什么
PHP-FPM 即 PHP-FastCGI Process Manager， 它是 FastCGI 的实现，并提供了进程管理的功能。进程包含 master 进程和 worker 进程两种；master 进程只有一个，负责监听端口，接收来自服务器的请求，而 worker 进程则一般有多个（具体数量根据实际需要进行配置），每个进程内部都会嵌入一个 PHP 解释器，是代码真正执行的地方。

## Nginx 与 php-fpm 通信机制
当我们访问一个网站（如 www.test.com）的时候，处理流程是这样的：

  www.test.com
        |
        |
      Nginx
        |
        |
路由到 www.test.com/index.php
        |
        |
加载 nginx 的 fast-cgi 模块
        |
        |
fast-cgi 监听 127.0.0.1:9000 地址
        |
        |
www.test.com/index.php 请求到达 127.0.0.1:9000
        |
        |
     等待处理...
## Nginx 与 php-fpm 的结合
在 Linux 上，nginx 与 php-fpm 的通信有 tcp socket 和 unix socket 两种方式。

tcp socket 的优点是可以跨服务器，当 nginx 和 php-fpm 不在同一台机器上时，只能使用这种方式。

Unix socket 又叫 IPC(inter-process communication 进程间通信) socket，用于实现同一主机上的进程间通信，这种方式需要在 nginx配置文件中填写 php-fpm 的 socket 文件位置。

两种方式的数据传输过程如下图所示：
![image](https://github.com/weilyf2017/Blog/blob/master/images/php-fpm%E9%80%9A%E4%BF%A1.png)

### 二者的不同：

由于 Unix socket 不需要经过网络协议栈，不需要打包拆包、计算校验和、维护序号和应答等，只是将应用层数据从一个进程拷贝到另一个进程。所以其效率比 tcp socket 的方式要高，可减少不必要的 tcp 开销。不过，unix socket 高并发时不稳定，连接数爆发时，会产生大量的长时缓存，在没有面向连接协议的支撑下，大数据包可能会直接出错不返回异常。而 tcp 这样的面向连接的协议，可以更好的保证通信的正确性和完整性。
