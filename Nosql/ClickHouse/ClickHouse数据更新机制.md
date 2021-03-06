# ClickHouse介绍

ClickHouse是一款用于大数据实时分析的列式数据库管理系统，而非数据库。通过向量化执行以及对cpu底层指令集（SIMD）的使用，它可以对海量数据进行并行处理，从而加快数据的处理速度。

### 主要优点有：

- 1）为了高效的使用CPU，数据不仅仅按列存储，同时还按向量进行处理；

- 2）数据压缩空间大，减少io；处理单查询高吞吐量每台服务器每秒最多数十亿行；

- 3）索引非B树结构，不需要满足最左原则；只要过滤条件在索引列中包含即可；即使在使用的数据不在索引中，由于各种并行处理机制ClickHouse全表扫描的速度也很快；

- 4）写入速度非常快，50-200M/s，对于大量的数据更新非常适用；

ClickHouse并非万能的，正因为ClickHouse处理速度快，所以也是需要为“快”付出代价。

### 缺点：

- 1）不支持事务，不支持真正的删除/更新；

- 2）不支持高并发，官方建议qps为100，可以通过修改配置文件增加连接数，但是在服务器足够好的情况下；

- 3）sql满足日常使用80%以上的语法，join写法比较特殊；最新版已支持类似sql的join，但性能不好；

- 4）尽量做1000条以上批量的写入，避免逐行insert或小批量的insert，update，delete操作，因为ClickHouse底层会不断的做异步的数据合并，会影响查询性能，这个在做实时数据写入的时候要尽量避开；

- 5）Clickhouse快是因为采用了并行处理机制，即使一个查询，也会用服务器一半的cpu去执行，所以ClickHouse不能支持高并发的使用场景，默认单查询使用cpu核数为服务器核数的一半，安装时会自动识别服务器核数，可以通过配置文件修改该参数；


针对数据高可用，我们对数据更新机制做了如下设计：

### 全量数据导入流程

![images](https://github.com/foxliang/Blog/blob/master/images/%E5%85%A8%E9%87%8F%E5%AF%BC%E5%85%A5%E6%95%B0%E6%8D%AE.jpeg)

全量数据的导入过程比较简单，仅需要将数据先导入到临时表中，导入完成之后，再通过对正式表和临时表进行ReName操作，将对数据的读取从老数据切换到新数据上来。

### 增量数据的导入过程

![images](https://github.com/foxliang/Blog/blob/master/images/%E5%A2%9E%E9%87%8F%E6%95%B0%E6%8D%AE%E7%9A%84%E5%AF%BC%E5%85%A5%E8%BF%87%E7%A8%8B.jpeg)

增量数据的导入过程，我们使用过两个版本。

由于ClickHouse的delete操作过于沉重，所以最早是通过删除指定分区，再把增量数据导入正式表的方式来实现的。

这种方式存在如下问题：一是在增量数据导入的过程中，数据的准确性是不可保证的，如果增量数据越多，数据不可用的时间就越长；二是ClickHouse删除分区的动作，是在接收到删除指令之后内异步执行，执行完成时间是未知的。如果增量数据导入后，删除指令也还在异步执行中，会导致增量数据也会被删除。最新版的更新日志说已修复这个问题。

针对以上情况，我们修改了增量数据的同步方案。在增量数据从Hive同步到ClickHouse的临时表之后，将正式表中数据反写到临时表中，然后通过ReName方法切换正式表和临时表。

通过以上流程，基本可以保证用户对数据的导入过程是无感知的。

参考:https://cloud.tencent.com/developer/article/1458155
