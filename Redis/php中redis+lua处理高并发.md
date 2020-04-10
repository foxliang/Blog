# 使用:
在php的redis中使用lua

```
    $redis->eval()函数

    eval($lua,$data,$num);

    $lua 要执行的lua命令 :

    $data传进去的参数(必须是数组):

    $num表示第二个参数数组中 有几个是参数(数组其他剩下来的是附加参数) 

    其中 lua中使用参数用的是 KEYS[1]  KEYS[2]  使用附加参数是 ARGV[1] 
```
php代码：
```
    public function unlock2($key){
        $redis = new Redis(); #实例化redis类
        $redis->connect('127.0.0.1'); #连接服务器
        $lua = <<<EOD
            local key = KEYS[1];
            local value = ARGV[1];
            if redis.call('get',key)==value then
                return redis.call('del',key)
            else
                return 0
            end
        EOD;
        $arr = [$key,$redis->get($key)];
        return $redis->eval($lua,$arr,1);
    }
```
# 优势:
Lua 嵌入 Redis 优势: 
减少网络开销: 不使用 Lua 的代码需要向 Redis 发送多次请求, 而脚本只需一次即可, 减少网络传输;
原子操作: Redis 将整个脚本作为一个原子执行, 无需担心并发, 也就无需事务;
复用: 脚本会永久保存 Redis 中, 其他客户端可继续使用.
# 不足之处：
该方案基于单个写节点的 Redis集群，无法适用于多个写节点的Redis集群；
Redis 执行 Lua 脚本 具有了原子性， 但是 Lua脚本内的 多个写操作 没有实现 原子性(事务)。
