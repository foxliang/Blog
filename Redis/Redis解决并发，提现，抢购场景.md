在秒杀，抢购等并发场景下，可能会出现超卖的现象，在PHP语言中并没有原生提供并发的解决方案，因此就需要借助其他方式来实现并发控制。

#### 列出常见的解决方案有：

1.使用队列，额外起一个进程处理队列，并发请求都放到队列中，由额外进程串行处理，并发问题就不存在了，但是要额外进程支持以及处理延迟严重，本文不先不讨论这种方法。

2.利用数据库事务特征，做原子更新，此方法需要依赖数据库的事务特性。

3.借助文件排他锁，在处理下单请求的时候，用flock锁定一个文件，成功拿到锁的才能处理订单。

##### 一、利用 Redis 事务特征

redis 事务是原子操作，可以保证订单处理的过程中数据没有被其它并发的进程修改。

可以用 watch 抓住redis的key

```
    $this->redis = new Redis();
    $this->redis->pconnect('127.0.0.1', 6379);
    $now_redis_amount = $this->redis->get($redis_id); //获取余额
    //Redis Watch 命令用于监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断
    $this->redis->watch($redis_id);
    $this->redis->multi();
    //修改redis
    $this->redis->set($redis_id, $set_redis_amount);
    $res = $this->redis->exec();
    if ($res === false) {
        return array('code' => 500, 'msg' => 'Update Redis Error');
    }
    return array('code' => 200, 'msg' => 'ok');
    
```

经测试 要比 MySQL 的select * for update  性能上要好的多
