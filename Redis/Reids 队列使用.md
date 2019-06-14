Reids是一个比较高级的开源key-value存储系统，采用ANSIC实现。其与Memcached类似，但是支持持久化数据存储。

入队操作


```
$redis = new Redis(); 
$redis->connect('127.0.0.1',6379); 

        $arr = [
                ['name' => 'xiaoming', 'age' => 20],
                ['name' => 'xiaoli', 'age' => 20],
                ['name' => 'sunzi', 'age' => 20],
                ['name' => 'mingwan', 'age' => 20],
                ['name' => 'lida', 'age' => 20],
                ['name' => 'kerong', 'age' => 20],
        ];
        foreach ($arr as $k => $v) {
            $this->redis->rpush("mylist", json_encode($v)); //加入队列值
        }
        echo '队列已经加入完成';
```

出队操作 


```
      $count = $this->redis->lSize('mylist'); //获取队列的长度
        var_dump($count);
        for ($i = 1; $i <= $count; $i++) {
            $val = $this->redis->LPOP('mylist');
            var_dump($val);
            echo "<br/>";
        }
```

用redis的list当作队列可能存在的问题
1)redis崩溃的时候队列功能失效

2)如果入队端一直在塞数据，而出队端没有消费数据，或者是入队的频率大而多，出队端的消费频率慢会导致内存暴涨

3)Redis的队列也可以像rabbitmq那样  即可以做消息的持久化，也可以不做消息的持久化。

当做持久话的时候，需要启动redis的dump数据的功能.暂时不建议开启持久化。
 

Redis其实只适合作为缓存，而不是数据库或是存储。它的持久化方式适用于救救急啥的，不太适合当作一个普通功能来用。应为dump时候，会影响性能，数据量小的时候还看不出来，当数据量达到百万级别，内存10g左右的时候，非常影响性能。
 

4)假如有多个消费者同时监听一个队列，其中一个出队了一个元素，另一个则获取不到该元素

5)Redis的队列应用场景是一对多或者一对一的关系，即有多个入队端，但是只有一个消费端(出队)
