当多个进程不在同一个系统中，就需要用分布式锁控制多个进程对资源的访问。
使用redis来实现分布式锁主要用到以下命令：


- SETNX KEY VALUE 如果key不存在，就设置key对应字符串value
- expire KEY seconds 设置key的过期时间
- del KEY 删除key


代码实现如下:

```
$redis = new Redis();
$redis->connect('127.0.0.1', 6379);
$ok = $redis->setNX($key, $value);
if ($ok) {
    //获取到锁
    ... do something ...
    $redis->del($key);
}
```

上面代码有没有问题呢？

如果我们在逻辑处理过程中出现了异常情况，导致KEY没有删除，那就出现了死锁了。所以一般我们在拿到锁之后再给KEY加一个过期时间
为了保证执行的原子性，使用了multi就有了如下代码


```
$redis->multi();
$redis->setNX($key, $value);
$redis->expire($key, $ttl);
$res = $redis->exec();
if($res[0]) {
    //获取到锁
    ... do something ...
    $redis->del($key);
}
```

但是这样的又有一个问题第一个请求成功了，之后的请求虽然没有拿到锁但是每次都刷新了锁的时间。这样我们设置锁过期时间的意义就不存在了。所以我们在拿到锁以后再进行过期时间的操作，这时候我们就可以祭出原子性操作的lua脚本，代码如下

```

$script = <<<EOT
    local key   = KEYS[1]
    local value = KEYS[2]
    local ttl   = KEYS[3]

    local ok = redis.call('setnx', key, value)

    if ok == 1 then
    redis.call('expire', key, ttl)
    end
    return ok
EOT;

$res = $redis->eval($script, [$key,$val, $ttl], 3);
if($res) {
    //获取到锁
    ... do something ...
    $redis->del($key);
}
```

借助lua脚本虽然解决了问题，但是未免有些麻烦，Redis从 2.6.12 版本开始， SET 命令的行为可以通过一系列参数来修改：

EX second ：设置键的过期时间为 second 秒。 SET key value EX second 效果等同于 SETEX key second value 。

PX millisecond ：设置键的过期时间为 millisecond 毫秒。 SET key value PX millisecond 效果等同于 PSETEX key millisecond value 。

NX ：只在键不存在时，才对键进行设置操作。 SET key value NX 效果等同于 SETNX key value 。

XX ：只在键已经存在时，才对键进行设置操作。

```

$ok = $redis->set($key, $random, array('nx', 'ex' => $ttl));

if ($ok) {
    //获取到锁
    ... do something ...
    if ($redis->get($key) == $random) {
        $redis->del($key);
    }
}
```

可以看到上面我们我们的值引入了一个随机数，这是为了防止逻辑处理时间过长导致锁的过期时间已经失效，这时候下一个请求就获得了锁，但是前一个请求在逻辑处理完直接删除了锁。

锁主要用在并发请求如秒杀等场景中，以上便是redis锁的实现。
