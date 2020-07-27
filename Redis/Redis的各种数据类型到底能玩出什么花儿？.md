# 正文

Redis作为一款NoSQL内存数据库，其丰富的数据类型、简单易用的命令、单机可达10万的高并发（官方数据），从面世以来就深受广大用户的喜爱。

Redis的五种数据类型，是我们学习Redis时的必修课，但是大多数人都只是去学它的命令、API，却不知道这些数据类型都能应用在哪些场景，那这些命令学起来也就会很快就忘，终究只是“纸上谈兵”。

用好这五种数据类型将给你的开发带来很大的便利，给你的程序带来很大的性能提升，同时这五种数据类型能玩出很多花样。

不过大多数同学，在实际的开发过程中，大多只用到了Redis五种数据类型中的1-3种，甚至有的只用过一种String类型。要么是业务场景简单用string足矣，要么就是根本不知道或想不到用别的数据类型更合适，那么即使是有些场景更适合用别的数据类型，可能自己也发觉不到。所以今天就来聊聊Redis的各种数据类型以及各自适用于什么场景。

其实这个问题也是面试中问到Redis时的一个“开篇词”，问这个问题主要有两个原因：
 
第一，看看你到底有没有全面的了解Redis有哪些功能，一般怎么来用，什么场景用什么数据类型，就怕你只会最简单的kv操作
 
第二，看看你在实际项目里都怎么玩儿过Redis，经验是否丰富

要是你回答的不好，没说出几种数据类型，也没说什么场景，你完了，面试官对你印象肯定不好，觉得你平时就是做个简单的set和get，没思考过。

废话不多说，进入正题吧。Redis一共提供了5种数据类型，分别是String，Hash，List，Set，sorted set(Zset)，下面就从各个数据类型的基本常用命令和使用场景分别说说吧。

## String字符串

String字符串结构的常用命令

#字符串常用操作
```
SET  key  value //存入字符串键值对
MSET  key  value [key value ...] //批量存储字符串键值对
SETNX  key  value //存入一个不存在的字符串键值对
GET  key //获取一个字符串键值
MGET  key  [key ...] //批量获取字符串键值
DEL  key  [key ...] //删除一个键
EXPIRE  key  seconds //设置一个键的过期时间(秒)
#原子加减
INCR  key //将key中储存的数字值加1
DECR  key //将key中储存的数字值减1
INCRBY  key  increment //将key所储存的值加上increment
DECRBY  key  decrement //将key所储存的值减去decrement
```
这里列出了一些String常用命令，我们看一下这些String类型的这些命令可以应用到哪些场景。

#### 应用场景

1、单值缓存

即最简单的key-value的set和get，比如缓存个标识，开关等
```
SET key value
GET key
```
2、对象缓存

除了单值缓存我们还可以用String类型缓存对象，如下两种方式：
```
#1
SET user:1  value(json串)
GET user:1
#2
MSET user:1:name 编程大道 user:1:sex 1
MGET user:1:name user:1:sex
```
第一种直接将对象转换成json串作为value存储到redis，这种获取对象就比较简单了，直接get key拿到value转成对象即可，但有个缺点就是如果你要是修改对象的某一个字段，也得把整个对象的json串拿出来反序列化成对象，这将带来不必要的网络开销(即便是redis存在内存中，但实际我们的应用服务器和redis是隔离的，网络传输的开销也不容小觑)，同样，频繁的序列化反序列化也将会带来不小的性能开销，如果对于性能要求比较高的系统来说这将是一个灾难。
而第二种存储对象的方式则对于这种频繁修改对象某一个字段的场景就比较友好了，每个字段与值都是一个kv对，修改直接set k v覆盖就好了，但是存储多个字段时就没那么容易了，好在有mset批量操作的命令，网络开销由多次变为1次。

3、分布式锁

如下setnx命令是set if not exit的缩写，意思就是这个key不存在时才执行set。多个线程执行这条命令时只有一个线程会执行成功，则视为拿到锁。然后拿到锁的线程执行业务操作，执行完毕删除这个锁，释放锁。
```
#setnx key value
SETNX  product:10001  true   //返回1代表获取锁成功
SETNX  product:10001  true   //返回0代表获取锁失败
//执行业务操作
DEL  product:10001  //执行完业务释放锁
```
上述方式存在问题：程序意外终止可能会导致锁没办法释放，造成死锁。可以使用如下命令，既设置分布式锁又设置了key的过期时间
```
SET product:10001 true  ex  10  nx  //防止程序意外终止导致死锁
```
分Redis布式锁的详细实现可以参考我之前写的Redis分布式锁实战

4、计数器
```
INCR article:readcount:{文章id}
GET article:readcount:{文章id}
```
基于Redis原子自增命令incr可以实现诸如计数器的功能，我们都知道公众号文章，微博，博客都有一个阅读量的概念，我们就可以用这个计数器来实现，而且性能很高。


5、Web集群session共享解决方案

系统集群部署情况下首先要考虑的问题就是session共享问题，我们可以通过将原本存储在内存中由tomcat管理的session转移到由Redis来存储，实现分布式session的功能。spring框架提供了session共享的解决方案，即spring session + redis实现分布式session。

6、分布式系统全局序列号

分布式系统中要保证全局序列号的唯一性，可以使用Redis来维护一个自增的序列。

通过如下命令从Redis获取自增ID：
```
#INCR是一个原子自增命令
INCR orderId
```
分布式系统环境下通过Redis保证ID的自增性和唯一性，通过该命令获取ID每次都要和Redis进行交互，如果业务量很大，那么这将会很频繁。

所以可以一次性获取一定量的ID保存在JVM内存中，用完了再从Redis获取。这样减少了频繁的网络开销，但是缺点是可能会丢失(浪费)一部分ID，因为获取后服务可能挂了还没用完的ID可能就浪费了（当然你可以使用一些手段去保证不浪费，但没必要，浪费一点也是无所谓的）。

如下，每次获取1000个
```
#redis批量生成序列号提升性
INCRBY  orderId  1000
```

## HASH结构
Hash常用操作
```
HSET key field value//存储一个哈希表key的键值
HSETNX key field value//存储一个不存在的哈希表key的键值
HMSET key field value [field value ...] //在一个哈希表key中存储多个键值对
HGET key field//获取哈希表key对应的field键值
HMGET key field [field ...]//批量获取哈希表key中多个field键值
HDEL key field [field ...]//删除哈希表key中的field键值
HLEN key//返回哈希表key中field的数量
HGETALL key//返回哈希表key中所有的键值
HINCRBY key field increment//为哈希表key中field键的值加上增量increment
```
应用场景

1、对象缓存

结合HASH结构的key-field-value的特性，类似于Java中的HashMap，内部也是“key-value”的形式，field刚好可以存对象的属性名，假设有如下数据，

我们可以用HMSET命令批量设置field-value，前面拼接用户的ID保证存多个用户的数据不会重复；HMGET批量获取field；MSET修改某一个field。
```
HMSET  achievement {userId}:name  小明 {userId}:score 89
HMSET  achievement 1:name  小明 1:score 89
HMSET  achievement 2:name  小华 2:score 92
HMGET  achievement 1:name  1:score
```
对象与HSAH的关系就变成了下图这样

2、电商购物车

以用户id为key，商品id为field，商品数量为value可以实现购物车的常规操作。

购物车操作：
```
#添加商品
hset cart:10001 50005 1
#给某一个商品增加数量
hincrby cart:10001 50005 1
#购物车中商品总个数
hlen cart:10001
#删除商品
hdel cart:10001 50005
#获取购物车所有商品
hgetall cart:10001
```
对应购物车的几个常用操作可以想象使用Redis如何实现

### Hash结构优缺点

优点

将同类数据归类整合储存（同一个key），方便数据管理

相比String操作，对内存与cpu的消耗更小

相比String储存更节省空间

缺点

过期功能不能使用在field上，只能用在key上

Redis集群架构下不适合大规模使用

## List结构

List常用操作

我们可以认为列表的左边叫头，右边叫尾

常用命令
```
LPUSH key value [value ...] //将一个或多个值value插入到key列表的表头(最左边)
```
应用场景

1、实现常见的数据结构

基于List的特性及丰富的命令可以实现常用的集中数据结构：

1）Stack(栈) = LPUSH + LPOP ，FILO先入后出

结合LPUSH和LPOP命令实现栈的先进后出的特性，LPUSH从左边入栈，LPOP从左边出栈，先进入的后出来。相当于入口出口是一个。

2）Queue(队列）= LPUSH + RPOP，FIFO先进先出

结合LPUSH和RPOP命令实现队列的先进先出的特性，LPUSH从左边入队，RPOP从右边出队，先进来的先出来。相当于入口出口各在两边。

3）Blocking MQ(阻塞队列）= LPUSH + BRPOP

结合LPUSH和BRPOP实现阻塞队列，BRPOP比RPOP多了一个timeout的参数，是一个等待的最大时间，如果在这个时间内拿不到数据则返回空。

2、微博消息和微信公号消息

例如，walking本人关注了人民网、华为中国、京港地铁等大V，假设人民网发了一条微博，ID为30033，我关注了他，那么就会往我的msg这个队列里push这个微博ID，我在打开我的微博时，就会从这个我专属的msg队列里取前几个微博ID展示给我看，所以这个就牵涉到了几个关键点：

1）人民网发了一条微博，ID为30033，消息ID入队
```
LPUSH  msg:{walking-ID}  30033
```
2）华为中国发微博，ID为30055，消息入队
```
LPUSH  msg:{walking-ID} 30055
```
3）我登录进去，会给我展示最新微博消息，那么就从我的消息队列里取最新的前5条显示在首页
```
LRANGE  msg:{walking-ID}  0  5
```

## SET结构

Set常用操作
```
SADD  key  member  [member ...]//往集合key中存入元素，元素存在则忽略，若key不存在则新建
SREM  key  member  [member ...]//从集合key中删除元素
SMEMBERS  key //获取集合key中所有元素
SCARD  key//获取集合key的元素个数
SISMEMBER  key  member//判断member元素是否存在于集合key中
SRANDMEMBER  key  [count]//从集合key中选出count个元素，元素不从key中删除
SPOP  key  [count]//从集合key中选出count个元素，元素从key中删除
```
set运算操作
```
SINTER key [key ...] //交集运算
SINTERSTORE destination key [key ..]//将交集结果存入新集合destination中
SUNION key [key ..] //并集运算
SUNIONSTORE destination key [key ...]//将并集结果存入新集合destination中
SDIFF key [key ...] //差集运算
SDIFFSTORE destination key [key ...]//将差集结果存入新集合destination中
```
应用场景

1、微信抽奖小程序

想必大家都用过微信里的抽奖小程序吧，如下图，我们可以点击立即参与进行抽奖，还可以查看所有参与人员，最后就是开奖的功能，一共三个关键点

我们看一下这三个关键点用set数据类型怎么实现：

1）点击参与抽奖，则将用户ID加入集合
```
SADD key {userlD}
```
2）查看参与抽奖所有用户
```
SMEMBERS key
```
3）抽取count名中奖者
```
SRANDMEMBER key [count]//返回但不从set中剔除
```
如果设置了一等奖二等奖三等奖...，并且每人只能得一种，则可以用SPOP key count

2、微信微博点赞，收藏，标签

比如walking发了一条朋友圈，有人点赞

1) 点赞 点赞就把点赞这个人的ID加到这个点赞的集合中
```
SADD  like:{消息ID}  {用户ID}
```
2) 取消点赞 从集合中移除用户ID
```
SREM like:{消息ID}  {用户ID}
```
3) 检查用户是否点过赞
```
SISMEMBER  like:{消息ID}  {用户ID}
```
4) 获取点赞的用户列表
```
SMEMBERS like:{消息ID}
```
5) 获取点赞用户数
```
SCARD like:{消息ID}
```
Set集合运算操作的应用场景

基于Redisset集合提供的丰富的命令，我们可以对集合轻松的实现交并差的运算。例如，现有集合set1，set12，set3，元素如下：
```
set1：{a,b,c}
```
对集合进行交、并、差的运算
```
SINTER set1 set2 set3 //交集--> { c } 
```
通过这些基本操作我们看可以实现什么样的业务需求。

3、集合操作实现社交软件关注模型

社交软件的用户关注模型，如QQ的好友，微博的关注，抖音、快手的关注、微信的公众号关注，这些社交软件都会做一个这样的功能，那就是用户关系的关注模型推荐，包括共同关注的人、可能认识的人、

首先看一下walking、chenmowanger、Hollis关注的人，如下：

1)walking关注的人:
```
walkingSet-->{chenmowanger, ImportNew, Hollis}
```
2) chenmowanger关注的人:
```
chenmowangerSet-->{walking, ImportNew, Hollis, JavaGuide}
```
3) Hollis关注的人:
```
HollisSet--> {waking, ImportNew, JavaGuide, feichao, CodeSheep}
```

每个人的关注列表都是一个Redis的set集合，然后当walking点到chenmowanger的主页，就会有个区域专门展示我和二哥的一些关注情况：

4) walking和chenmowanger共同关注:

也就是看哪些人在我的集合里也在二哥的集合里

//两个集合求并集
```
SINTER walkingSet zhangyixingSet--> {ImportNew, Hollis}
```
5) 我关注的人也关注他(chenmowanger):

看我关注的人的关注列表里是不是有某个人，比如我进入chenmowanger的主页，可以展示我关注的人里还有谁也关注了chenmowanger
```
SISMEMBER ImportNewSet chenmowanger
```
6) 我可能认识的人:
求差集，以前面这个集合为准，看二哥关注的那些人有哪些我还没关注，于是我就赶紧关注了JavaGuide（Guide哥）
```
SDIFF chenmowangerSet walkingSet->{walking, JavaGuide}
```
4、集合操作实现电商商品筛选
先看一下这个图是不是很熟悉，选购手机时，有一个筛选的功能

如上图，电商网站买手机，进到这个页面根据各种条件搜手机，我们想一想用Redis如何实现呢？（当然了，这里并不是说人家就完全用Redis实现这一套搜索，其实主要还是用搜索引擎那些中间件，这里只是说明可以用Redis实现~）
在上架商品时维护商品，添加商品的同时把对应的商品添加到对应的set集合里即可，如下举例
```
//品牌-华为
SADD  brand:huawei  P30 Mate30 荣耀Play4 nova7
//品牌-小米
SADD  brand:xiaomi  mi6 mi8 mi9 mi10
//品牌-iPhone
SADD  brand:iPhone iphone8 iphone8plus iphoneX iphone11
//操作系统-Android
SADD os:android  P30 Mate30 荣耀Play4 nova7 mi6 mi8 mi9 mi10
//CPU品牌-骁龙
SADD cpu:brand:xiaolong iphone8 iphone8plus iphoneX iphone11 mi6 mi8 mi9 mi10
//CPU品牌-麒麟
SADD cpu:brand:qilin  P30 Mate30 荣耀Play4 nova7
//运行内存-8G
SADD ram:8G P30 Mate30 荣耀Play4 nova7 mi6 mi8 mi9 mi10 iphone8 iphone8plus iphoneX iphone11
//多条件查询 操作系统Android，CPU品牌骁龙，运行内存8G
SINTER  os:android cpu:brand:xiaolong  ram:8G -->{mi6 mi8 mi9 mi10}
```

假设我们维护了各种品牌，手机所属的操作系统，CPU品牌，运行内存等，那么我们在勾选条件查找时就可以用勾选的各个集合求他的交集就行了。

## ZSet有序集合

zset是有序的set集合，通过传入的分值进行排序

ZSet常用操作
```
ZADD key score member [[score member]…]//往有序集合key中加入带分值元素
ZREM key member [member …]  //从有序集合key中删除元素
ZSCORE key member //返回有序集合key中元素member的分值
ZINCRBY key increment member//为有序集合key中元素member的分值加上increment 
ZCARD key//返回有序集合key中元素个数
ZRANGE key start stop [WITHSCORES]//正序获取有序集合key从start下标到stop下标的元素
ZREVRANGE key start stop [WITHSCORES]//倒序获取有序集合key从start下标到stop下标的元素
```
Zset集合操作
```
ZUNIONSTORE destkey numkeys key [key ...] //并集计算 
ZINTERSTORE destkey numkeys key [key …]//交集计算
```

应用场景

1、Zset集合操作实现排行榜

我们都知道微博热点，新闻热榜，投票排行榜等都有一个排名的概念，如下图百度热榜，展示的是实时的点击量比较高的新闻（假设这些新闻的ID为1001-1010），每个新闻都有一个热点值，一般按点击量，1001这个新闻热点是484W，1002这个是467W，实时的，可能等会再看就不一样了，那么我们看下用Redis咋实现。

1）点击新闻

每次有人点击这个新闻，那么久ius给他的分值加1
```
ZINCRBY  hotNews:20200722  1  1001 //新闻ID为1001的新闻分值加一
```
2）展示当日排行前十

取集合中的前10个元素
```
ZREVRANGE  hotNews:20200722  0  10  WITHSCORES
```
3）七日热点榜单计算
```
ZUNIONSTORE  hotNews:20200715-20200721  7 hotNews:20200715 hotNews:20200716... hotNews:20200721
```
4）展示七日排行前十
```
ZREVRANGE hotNews:20190813-20190819  0  10  WITHSCORES
```
更多应用场景


微信<摇一摇><抢红包>


滴滴打车、摩拜单车<附近的车>


美团和饿了么<附近的餐馆>


搜索自动补全


布隆过滤器


参考：

https://www.bilibili.com/video/BV1if4y1R7ns

https://mp.weixin.qq.com/s/ZSQ9vCkWXYuLrKS0UJ4RIg

https://juejin.im/post/5f1cfda45188252e4a280773
