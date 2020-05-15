CentOS 版本：centos-release-8.1-1.1911.0.9.el8.x86_64

Nginx 版本：nginx/1.14.1

Nginx 介绍
Nginx 是开源、高性能、高可靠的 Web 和反向代理服务器，而且支持热部署，几乎可以做到 7 * 24 小时不间断运行，即使运行几个月也不需要重新启动，还能在不间断服务的情况下对软件版本进行热更新。性能是 Nginx 最重要的考量，其占用内存少、并发能力强、能支持高达 5w 个并发连接数，最重要的是，Nginx 是免费的并可以商业化，配置使用也比较简单。

Nginx 的最重要的几个使用场景：

静态资源服务，通过本地文件系统提供服务；

反向代理服务，延伸出包括缓存、负载均衡等；

API 服务，OpenResty ；

对于前端来说 Node.js 不陌生了，Nginx 和 Node.js 的很多理念类似，HTTP 服务器、事件驱动、异步非阻塞等，且 Nginx 的大部分功能使用 Node.js 也可以实现，但 Nginx 和 Node.js 并不冲突，都有自己擅长的领域。Nginx 擅长于底层服务器端资源的处理（静态资源处理转发、反向代理，负载均衡等），Node.js 更擅长上层具体业务逻辑的处理，两者可以完美组合，共同助力前端开发。

下面我们着重学习一下 Nginx 的使用。

先下载centos
```
docker pull centos
```
我自己做好了镜像 可以直接下载


运行
```
docker run -itd -p 8900:80 -p 8901:8080  --privileged --name centos centos /usr/sbin/init
```
加粗的内容要特别注意，不能遗忘

原因就是： 默认情况下，在第一步执行的是 /bin/bash，而因为docker中的bug，无法使用systemctl 

所以我们使用了 /usr/sbin/init 同时 --privileged 这样就能够使用systemctl了，但覆盖了默认的 /bin/bash


进入容器
```
docker exec -it f870fe771dc4 /bin/bash
```
下载nginx
```
yum install nginx
```
来安装 Nginx，然后我们在命令行中 nginx -v 就可以看到具体的 Nginx 版本信息，也就安装完毕了。



安装之后开启 Nginx，如果系统开启了防火墙，那么需要设置一下在防火墙中加入需要开放的端口，下面列举几个常用的防火墙操作（没开启的话不用管这个）：
```
systemctl start firewalld  # 开启防火墙
systemctl stop firewalld   # 关闭防火墙
systemctl status firewalld # 查看防火墙开启状态，显示running则是正在运行
firewall-cmd --reload      # 重启防火墙，永久打开端口需要reload一下
# 添加开启端口，--permanent表示永久打开，不加是临时打开重启之后失效

firewall-cmd --permanent --zone=public --add-port=8888/tcp
```
# 查看防火墙，添加的端口也可以看到
```
firewall-cmd --list-all
```
然后设置 Nginx 的开机启动：
```
systemctl enable nginx
```
启动 Nginx （其他命令后面有详细讲解）：
```
systemctl start nginx
```
然后访问你的 IP，这时候就可以看到 Nginx 的欢迎页面了～ Welcome to nginx！ 👏



 

配置反向代理
反向代理是工作中最常用的服务器功能，经常被用来解决跨域问题，下面简单介绍一下如何实现反向代理。

首先进入 Nginx 的主配置文件：
```
vim /etc/nginx/nginx.conf
```

```
http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    keepalive_timeout   65;
    types_hash_max_size 2048;

    include             /etc/nginx/mime.types;
    default_type        application/octet-stream;

    # Load modular configuration files from the /etc/nginx/conf.d directory.
    # See http://nginx.org/en/docs/ngx_core_module.html#include
    # for more information.
    include /etc/nginx/conf.d/*.conf;

    server {
        listen       80 default_server;
        listen       [::]:80 default_server;
        server_name  _;
        root         /usr/share/nginx/html;

        # Load configuration files for the default server block.
        include /etc/nginx/default.d/*.conf;

        location / {
         proxy_pass http://www.bilibili.com;
        }

        error_page 404 /404.html;
            location = /40x.html {
        }

        error_page 500 502 503 504 /50x.html;
            location = /50x.html {
        }
    }
    server {
        listen      8080 default_server;
        listen      [::]:8080 default_server;
        server_name _;
        root        /usr/share/nginx/html;

        include /etc/nginx/defalut.d/*.conf;

        location / {
        }
        error_page 404 /404.html;
              location = /40x.html {
        }

        error_page 500 502 503 504 /50x.html;
              location = /50x.html {
        }
    }
````
端口转换

80--->http://www.bilibili.com; 8080--->80
