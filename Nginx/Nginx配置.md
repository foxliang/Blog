# 常用配置项
在工作中，我们与 Nginx 打交道更多的是通过其配置文件来进行。那么掌握这些配置项各自的作用就很有必要了。

首先，nginx.conf 的内容通常是这样的：

```
...              
...            #核心摸块

events {        #事件模块
 
   ...
}

http {     # http 模块

    server {      # server块
     
        location [PATTERN] {  # location块
        
            ...
        }
        location [PATTERN] {
        
            ...
        }
    }
    server {
      ...
    }
    
}

mail {     # mail 模块
     
     server {    # server块
          ...
    }

}

```
我们依次看一下每个模块一般有哪些配置项：

### 核心模块
user admin; #配置用户或者组。

worker_processes 4; #允许生成的进程数，默认为1 

pid /nginx/pid/nginx.pid; #指定 nginx 进程运行文件存放地址 

error_log log/error.log debug; #错误日志路径，级别。
### 事件模块
```
events { 
    accept_mutex on; #设置网路连接序列化，防止惊群现象发生，默认为on 
    
    multi_accept on; #设置一个进程是否同时接受多个网络连接，默认为off 
    
    use epoll; #事件驱动模型select|poll|kqueue|epoll|resig
    
    worker_connections 1024; #最大连接数，默认为512
}
```
http 模块
```
http {
    include       mime.types;   #文件扩展名与文件类型映射表
    
    default_type  application/octet-stream; #默认文件类型，默认为text/plain
    
    access_log off; #取消服务日志    

    sendfile on;   #允许 sendfile 方式传输文件，默认为off，可以在http块，server块，location块。
    
    sendfile_max_chunk 100k;  #每个进程每次调用传输数量不能大于设定的值，默认为0，即不设上限。
    
    keepalive_timeout 65;  #连接超时时间，默认为75s，可以在http，server，location块。

    server 
    {
            keepalive_requests 120; #单连接请求上限次数。
            
            listen 80; #监听端口
            
            server_name  127.0.0.1;   #监听地址      
            
            index index.html index.htm index.php;
            
            root your_path;  #根目录
          
            location ~ \.php$
            {
                  fastcgi_pass unix:/var/run/php/php7.1-fpm.sock;
                  
                  #fastcgi_pass 127.0.0.1:9000;
                  
                  fastcgi_index index.php;
                  
                  include fastcgi_params;
            }

    }
}
```
配置项解析
worker_processes

worker_processes 用来设置 Nginx 服务的进程数。该值推荐使用 CPU 内核数。

worker_cpu_affinity

worker_cpu_affinity 用来为每个进程分配CPU的工作内核，参数有多个二进制值表示，每一组代表一个进程，每组中的每一位代表该进程使用CPU的情况，1代表使用，0代表不使用。所以我们使用 worker_cpu_affinity 0001 0010 0100 1000;来让进程分别绑定不同的核上。默认情况下worker进程不绑定在任何一个CPU上。

worker_rlimit_nofile

设置毎个进程的最大文件打开数。如果不设的话上限就是系统的 ulimit –n的数字，一般为65535。

worker_connections

设置一个进程理论允许的最大连接数，理论上越大越好，但不可以超过 worker_rlimit_nofile 的值。

use epoll

设置事件驱动模型使用 epoll。epoll 是 Nginx 支持的高性能事件驱动库之一。是公认的非 常优秀的事件驱动模型。

accept_mutex off

关闭网络连接序列化，当其设置为开启的时候，将会对多个 Nginx 进程接受连接进行序列化，防止多个进程对连接的争抢。当服务器连接数不多时，开启这个参数会让负载有一定程度的降低。但是当服务器的吞吐量很大时，为了效率，请关闭这个参数；并且关闭这个参数的时候也可以让请求在多个 worker 间的分配更均衡。所以我们设置 accept_mutex off;

multi_accept on

设置一个进程可同时接受多个网络连接

Sendfile on

Sendfile是 Linux2.0 以后的推出的一个系统调用,它能简化网络传输过程中的步骤，提高服务器性能。

不用 sendfile的传统网络传输过程：

硬盘 >> kernel buffer >> user buffer >> kernel socket buffer >> 协议栈

用 sendfile()来进行网络传输的过程：

硬盘 >> kernel buffer (快速拷贝到 kernelsocket buffer) >>协议栈

tcp_nopush on;

设置数据包会累积一下再一起传输，可以提高一些传输效率。 tcp_nopush 必须和 sendfile 搭配使用。

tcp_nodelay on;

小的数据包不等待直接传输。默认为on。 看上去是和 tcp_nopush 相反的功能，但是两边都为 on 时 nginx 也可以平衡这两个功能的使用。

keepalive_timeout

HTTP 连接的持续时间。设的太长会使无用的线程变的太多。这个根据服务器访问数量、处理速度以及网络状况方面考虑。

send_timeout

设置 Nginx 服务器响应客户端的超时时间，这个超时时间只针对两个客户端和服务器建立连接后，某次活动之间的时间，如果这个时间后，客户端没有任何活动，Nginx服务器将关闭连接

gzip on

启用 gzip，对响应数据进行在线实时压缩,减少数据传输量。

gzip_disable "msie6"

Nginx服务器在响应这些种类的客户端请求时，不使用 Gzip 功能缓存应用数据，gzip_disable “msie6”对IE6浏览器的数据不进行 GZIP 压缩。

常用的配置项大致这些，对于不同的业务场景，有的需要额外的其他配置项，这里不做展开。

### 其他
http 配置里有 location 这一项，它是用来根据请求中的 uri 来为其匹配相应的处理规则。

```
location 查找规则

location  = / {
  # 精确匹配 / ，主机名后面不能带任何字符串
  [ config A ]
}

location  / {
  # 因为所有的地址都以 / 开头，所以这条规则将匹配到所有请求
  # 但是正则和最长字符串会优先匹配
  [ config B ]
}

location /documents/ {
  # 匹配任何以 /documents/ 开头的地址，匹配符合以后，还要继续往下搜索
  # 只有后面的正则表达式没有匹配到时，这一条才会采用这一条
  [ config C ]
}

location ~ /documents/Abc {
  # 匹配任何以 /documents/Abc 开头的地址，匹配符合以后，还要继续往下搜索
  # 只有后面的正则表达式没有匹配到时，这一条才会采用这一条
  [ config CC ]
}

location ^~ /images/ {
  # 匹配任何以 /images/ 开头的地址，匹配符合以后，停止往下搜索正则，采用这一条。
  [ config D ]
}

location ~* \.(gif|jpg|jpeg)$ {
  # 匹配所有以 gif,jpg或jpeg 结尾的请求
  # 然而，所有请求 /images/ 下的图片会被 config D 处理，因为 ^~ 到达不了这一条正则
  [ config E ]
}

location /images/ {
  # 字符匹配到 /images/，继续往下，会发现 ^~ 存在
  [ config F ]
}

location /images/abc {
  # 最长字符匹配到 /images/abc，继续往下，会发现 ^~ 存在
  # F与G的放置顺序是没有关系的
  [ config G ]
}

location ~ /images/abc/ {
  # 只有去掉 config D 才有效：先最长匹配 config G 开头的地址，继续往下搜索，匹配到这一条正则，采用
    [ config H ]
}

```
正则查找优先级从高到低依次如下：

“ = ” 开头表示精确匹配，如 A 中只匹配根目录结尾的请求，后面不能带任何字符串。

“ ^~ ” 开头表示uri以某个常规字符串开头，不是正则匹配

“ ~ ” 开头表示区分大小写的正则匹配;

“ ~* ”开头表示不区分大小写的正则匹配

“ / ” 通用匹配, 如果没有其它匹配,任何请求都会匹配到

负载均衡配置
Nginx 的负载均衡需要用到 upstream 模块，可通过以下配置来实现：

```
upstream test-upstream {
    ip_hash; # 使用 ip_hash 算法分配
 
    server 192.168.1.1; # 要分配的 ip
    server 192.168.1.2;
}

server {

    location / {       
        proxy_pass http://test-upstream;
    }
    
}
```
上面的例子定义了一个 test-upstream 的负载均衡配置，通过 proxy_pass 反向代理指令将请求转发给该模块进行分配处理。
