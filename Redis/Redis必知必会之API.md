一些废话
- 1.Redis是一个开源、基于键值的存储系统、多种数据结构、功能丰富。
- 2.Redis支持持久化，断电不丢数据，对数据的更新异步保存到磁盘上。
- 3.Redis支持字符串、哈希、列表、集合、有序集合、位图、超小内存唯一计数、地理信息定位。
- 4.Redis支持多语言客户端，支持发布订阅，Lua脚本，事物，不依赖外部库，单线程模型，支持主从复制，高可用，分布式。
- 5.Redis典型使用场景有缓存系统、计数器，消息队列系统、排行榜、社交网络、实时系统。
- 6.启动方式分为

最简启动 redis-server
动态参数 redis-server --port 6380
配置文件 redis-server configPath
- 7.验证是否启动

ps -ef | grep redis
netstat -antpl | grep redis
redis-cli -h ip -p port ping
- 8.Redis客户端返回值有状态回复、错误回复、整数回复、字符串回复、多行字符串回复。
- 9.常用配置

daemonize 是否是守护进程
port 对外端口
logfile 日志
dir 工作目录
- 10.一次只运行一条命令，拒绝长慢命令，不要轻易执行 keys、flushdb、flushall、show lua script等。

- 11.其实redis不是单线程 例如如下API fysnc file descriptor、close file descriptor

## Redis API
1.通用命令

```
keys: 计算所有的键 O(n)
dbsize: 数据库大小
exists keys: key是否存在
del：删除key
expire key seconds: 设置过期时间
type key: 获取key的数据类型
ttl key: 查看key的剩余过期时间
persist key: 去掉key的过期时间
```
2.列表类型

```
rpush key value1 value2 valueN O(1-n)
lpush key value1 value2 valueN O(1-n)
linsert key before|after value newValue O(n)
lpop key
rpop key
lrem key count value (删除count个value元素) 0(n)
ltrim key start end (按照索引范围修剪列表) o(n)
lrange key start end (包含end) (获取列表制定索引范围)
lindex key index o(n)
llen key
lset key index newValue
blpop key timeout (lpop的阻塞版本)
brpop key timeout (rpop的阻塞版本)
lpush + lpop = stack
lpush + rpop = queue
lpush + ltrim = Capped Collection
lpush + brpop = Message Queue
```
3.字符串类型

场景：缓存、计数器、分布式锁
```
get key
set key value
incr key
decr key
incrby key k
decrby key k
setnx key value 值不存在才设置
set key value xx 存在才设置
mget 批量获取 O(n)
mset 批量设置 O(n)
getset key newvalue 设置新值返回旧值
append key value 将新值追加到旧值
strlen 获取值的长度
incrbyfloat key 3.5 增加对应key 3.5
getrange key start end
setrange key start value
```
4.集合类型

无序 无重复 支持集合间操作
```
sadd key element （添加）
srem key element （删除）
scard key
sismember key element
srandmember key count (随机选出count个元素)
spop key (随机弹出一个元素)
smembers key (取出所有元素 小心使用)
sscan (遍历集合)
sdiff 差集
sinter 交集
sunion 并集
sadd = 打标签
spop/srandmember = 随机
sadd + sinter = Social Graph
```
5.有序集合类型

```
zadd key score element O(logN)
zrem key element （删除）
zscore key element
zincrby key increScore element （增加分数）
zcard key (返回个数)
zrange key start end withscores (获取元素) （O(logN + m)）
zrangebyscore key minScore maxScore
zcount key minScore maxScore （O(logN + m)）
zremrangebyrank key start end (删除指定排名内的升序元素)
zremrangebyscore key start end (删除指定分数内的升序元素)
zrevrank
zrevrange
zrevrangebyscore
zinterstore
zunionstore
```
6.哈希类型

```
hget key filed
hset key field value
hdel key field
hgetall key O(n)
hexists key field
hlen key 获取字段数量
hmget key field1 field2 O(n)
hmset key field1 value1 field2 value2 O(n)
hincrby key field value
hvals key 返回hash key对应所有field的value O(n)
hkeys key 返回hash key对应的所有field O(n)
hsetnx key field value
hincrby key field intCounter
hincrbyfloat key field floatCounter
```
