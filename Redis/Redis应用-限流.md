## 在高并发场景下有三把利器保护系统：缓存、降级、和限流。

##### 缓存的目的是提升系统的访问你速度和增大系统能处理的容量；

##### 降级是当服务出问题或影响到核心流程的性能则需要暂时屏蔽掉。

而有些场景则需要限制并发请求量，如秒杀、抢购、发帖、评论、恶意爬虫等。

##### 限流算法

常见的限流算法有：计数器，漏桶、令牌桶。

###### 计数器

顾名思义就是来一个记一个，然后判断在有限时间窗口内的数量是否超过限制即可


```

function isActionAllowed($userId, $action, $period, $maxCount) 
{
    $redis = new Redis();
    $redis->connect('127.0.0.1', 6379);
    $key = sprintf('hist:%s:%s', $userId, $action);
    $now = msectime();   # 毫秒时间戳

    $pipe=$redis->multi(Redis::PIPELINE); //使用管道提升性能
    $pipe->zadd($key, $now, $now); //value 和 score 都使用毫秒时间戳
    $pipe->zremrangebyscore($key, 0, $now - $period); //移除时间窗口之前的行为记录，剩下的都是时间窗口内的
    $pipe->zcard($key);  //获取窗口内的行为数量
    $pipe->expire($key, $period + 1);  //多加一秒过期时间
    $replies = $pipe->exec();
    return $replies[2] <= $maxCount;
}
for ($i=0; $i<20; $i++){
    var_dump(isActionAllowed("110", "reply", 60*1000, 5)); //执行可以发现只有前5次是通过的
}

//返回当前的毫秒时间戳
function msectime() {
    list($msec, $sec) = explode(' ', microtime());
    $msectime = (float)sprintf('%.0f', (floatval($msec) + floatval($sec)) * 1000);
    return $msectime;
 }
```

###### 漏桶

漏桶(Leaky Bucket)算法思路很简单,水(请求)先进入到漏桶里,漏桶以一定的速度出水(接口有响应速率),当水流入速度过大会直接溢出(访问频率超过接口响应速率),然后就拒绝请求,可以看出漏桶算法能强行限制数据的传输速率:

具体代码实现如下

```

<?php

class Funnel {

    private $capacity;
    private $leakingRate;
    private $leftQuote;
    private $leakingTs;

    public function __construct($capacity, $leakingRate)
    {
        $this->capacity = $capacity;    //漏斗容量
        $this->leakingRate = $leakingRate;//漏斗流水速率
        $this->leftQuote = $capacity; //漏斗剩余空间
        $this->leakingTs = time(); //上一次漏水时间
    }

    public function makeSpace()
    {
        $now = time();
        $deltaTs = $now-$this->leakingTs; //距离上一次漏水过去了多久
        $deltaQuota = $deltaTs * $this->leakingRate; //可腾出的空间
        if($deltaQuota < 1) {  
            return;
        }
        $this->leftQuote += $deltaQuota;   //增加剩余空间
        $this->leakingTs = time();         //记录漏水时间
        if($this->leftQuota > $this->capacaty){
            $this->leftQuote - $this->capacity;
        }
    }

    public function watering($quota)
    {
        $this->makeSpace(); //漏水操作
        if($this->leftQuote >= $quota) {
            $this->leftQuote -= $quota;
            return true;
        }
        return false;
    }
}


$funnels = [];
global $funnel;

function isActionAllowed($userId, $action, $capacity, $leakingRate)
{
    $key = sprintf("%s:%s", $userId, $action);
    $funnel = $GLOBALS['funnel'][$key] ?? '';
    if (!$funnel) {
        $funnel  = new Funnel($capacity, $leakingRate);
        $GLOBALS['funnel'][$key] = $funnel;
    }
    return $funnel->watering(1);
}

for ($i=0; $i<20; $i++){
    var_dump(isActionAllowed("110", "reply", 15, 0.5)); //执行可以发现只有前15次是通过的
}
```

核心逻辑就是makeSpace，在每次灌水前调用以触发漏水，给漏斗腾出空间。

funnels我们可以利用Redis中的hash结构来存储对应字段，灌水时将字段取出进行逻辑运算后再存入hash结构中即可完成一次行为频度的检测。但这有个问题就是整个过程的原子性无法保证，意味着要用锁来控制，但如果加锁失败，就要重试或者放弃，这回导致性能下降和影响用户体验，同时代码复杂度也升高了，此时Redis提供了一个插件，Redis-Cell出现了。

###### 令牌桶

令牌桶算法(Token Bucket)和 Leaky Bucket 效果一样但方向相反的算法,更加容易理解.随着时间流逝,系统会按恒定1/QPS时间间隔(如果QPS=100,则间隔是10ms)往桶里加入Token(想象和漏洞漏水相反,有个水龙头在不断的加水),如果桶已经满了就不再加了.新请求来临时,会各自拿走一个Token,如果没有Token可拿了就阻塞或者拒绝服务.

令牌桶的另外一个好处是可以方便的改变速度.
一旦需要提高速率,则按需提高放入桶中的令牌的速率.
一般会定时(比如100毫秒)往桶中增加一定数量的令牌, 有些变种算法则实时的计算应该增加的令牌的数量.
具体实现可参考php 基于redis使用令牌桶算法实现流量控制

```
<?php
class TrafficShaper
{ 
    private $_config; // redis设定
    private $_redis;  // redis对象
    private $_queue;  // 令牌桶
    private $_max;    // 最大令牌数

    /**
     * 初始化
     * @param Array $config redis连接设定
     */
    public function __construct($config, $queue, $max)
    {
        $this->_config = $config;
        $this->_queue = $queue;
        $this->_max = $max;
        $this->_redis = $this->connect();
    }

    /**
     * 加入令牌
     * @param Int $num 加入的令牌数量
     * @return Int 加入的数量
     */
    public function add($num = 0)
    {
        // 当前剩余令牌数
        $curnum = intval($this->_redis->lSize($this->_queue));
        // 最大令牌数
        $maxnum = intval($this->_max);
        // 计算最大可加入的令牌数量，不能超过最大令牌数
        $num = $maxnum >= $curnum + $num ? $num : $maxnum - $curnum;
        // 加入令牌
        if ($num > 0) {
            $token = array_fill(0, $num, 1);
            $this->_redis->lPush($this->_queue, ...$token);
            return $num;
        }
        return 0;
    }

    /**
     * 获取令牌
     * @return Boolean
     */
    public function get()
    {
        return $this->_redis->rPop($this->_queue) ? true : false;
    }

    /**
     * 重设令牌桶，填满令牌
     */
    public function reset()
    {
        $this->_redis->delete($this->_queue);
        $this->add($this->_max);
    }

    private function connect()
    {
        try {
            $redis = new Redis();
            $redis->connect($this->_config['host'], $this->_config['port'], $this->_config['timeout'], $this->_config['reserved'], $this->_config['retry_interval']);
            if (empty($this->_config['auth'])) {
                $redis->auth($this->_config['auth']);
            }
            $redis->select($this->_config['index']);
        } catch (\RedisException $e) {
            throw new Exception($e->getMessage());
            return false;
        }
        return $redis;
    }
} 

$config = array(
    'host' => 'localhost',
    'port' => 6379,
    'index' => 0,
    'auth' => '',
    'timeout' => 1,
    'reserved' => NULL,
    'retry_interval' => 100,
);
// 令牌桶容器
$queue = 'mycontainer';
 // 最大令牌数
$max = 5;
// 创建TrafficShaper对象
$oTrafficShaper = new TrafficShaper($config, $queue, $max);
// 重设令牌桶，填满令牌
$oTrafficShaper->reset();
// 循环获取令牌，令牌桶内只有5个令牌，因此最后3次获取失败
for ($i = 0; $i < 8; $i++) {
    var_dump($oTrafficShaper->get());
}
// 加入10个令牌，最大令牌为5，因此只能加入5个
$add_num = $oTrafficShaper->add(10);
var_dump($add_num);
// 循环获取令牌，令牌桶内只有5个令牌，因此最后1次获取失败
for ($i = 0; $i < 6; $i++) {
    var_dump($oTrafficShaper->get());
}
?>
```

