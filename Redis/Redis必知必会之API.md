一些废话
1.Redis是一个开源、基于键值的存储系统、多种数据结构、功能丰富。
2.Redis支持持久化，断电不丢数据，对数据的更新异步保存到磁盘上。
3.Redis支持字符串、哈希、列表、集合、有序集合、位图、超小内存唯一计数、地理信息定位。
4.Redis支持多语言客户端，支持发布订阅，Lua脚本，事物，不依赖外部库，单线程模型，支持主从复制，高可用，分布式。
5.Redis典型使用场景有缓存系统、计数器，消息队列系统、排行榜、社交网络、实时系统。
6.启动方式分为

最简启动 redis-server
动态参数 redis-server --port 6380
配置文件 redis-server configPath
7.验证是否启动

ps -ef | grep redis
netstat -antpl | grep redis
redis-cli -h ip -p port ping
8.Redis客户端返回值有状态回复、错误回复、整数回复、字符串回复、多行字符串回复。
9.常用配置

daemonize 是否是守护进程
port 对外端口
logfile 日志
dir 工作目录
10.一次只运行一条命令，拒绝长慢命令，不要轻易执行 keys、flushdb、flushall、show lua script等。

11.其实redis不是单线程 例如如下API fysnc file descriptor、close file descriptor
