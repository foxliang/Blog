## 下载镜像
```
docker pull wurstmeister/zookeeper  
docker pull wurstmeister/kafka  
```
## 启动zookeeper容器

```
docker run -d --name zookeeper -p 2181:2181 -t wurstmeister/zookeeper
```

### 启动kafka容器

```
docker run --name kafka -p 9092:9092  --link zookeeper:zookeeper    -e KAFKA_ADVERTISED_HOST_NAME=localhost  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -t wurstmeister/kafka```
```

- 这里连接的时候可能会出现超时情况

可以修改ZooKeeper的连接超时时间
```
vi opt/kafka_2.13-2.6.0//config/server.properties

zookeeper.connection.timeout.ms =600000 
```

## 测试kafka 进入kafka容器的命令行

```
docker exec -it kafka /bin/bash
```
进入kafka所在目录

```
cd opt/kafka_2.13-2.6.0
```
启动消息发送方

```
 ./bin/kafka-console-producer.sh --broker-list localhost:9092 --topic mykafka
```
克隆会话 进入kafka所在目录

```
cd opt/kafka_2.13-2.6.0
```
启动消息接收方

```
   ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic mykafka --from-beginning
```
在消息发送方输入123456

在消息接收方查看

如果看到123456 消息发送完成

## 集群搭建
使用docker命令可快速在同一台机器搭建多个kafka，只需要改变brokerId和端口

```
docker run -d --name kafka1 -p 9093:9093 -e KAFKA_BROKER_ID=1 -e KAFKA_ZOOKEEPER_CONNECT=192.168.1.100:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.1.100:9093 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9093 -t wurstmeister/kafka
```
#### 创建Replication为2，Partition为2的topic 在kafka容器中的opt/kafka_2.12-1.1.0/目录下输入

```
bin/kafka-topics.sh --create --zookeeper 192.168.1.100:2181 --replication-factor 2 --partitions 2 --topic partopic
```
#### 查看topic的状态 在kafka容器中的opt/kafka_2.12-1.1.0/目录下输入

```
bin/kafka-topics.sh --describe --zookeeper 192.168.1.100:2181 --topic partopic
```
输出结果：

```
Topic:partopic  PartitionCount:2    ReplicationFactor:2 Configs:
    Topic: partopic Partition: 0    Leader: 0   Replicas: 0,1   Isr: 0,1
    Topic: partopic Partition: 1    Leader: 0   Replicas: 1,0   Isr: 0,1
```
显示每个分区的Leader机器为broker0，在broker0和1上具有备份，Isr代表存活的备份机器中存活的。 当停掉kafka1后，

```
docker stop kafka1
```
再查看topic状态，输出结果：

```
Topic:partopic  PartitionCount:2    ReplicationFactor:2 Configs:
    Topic: partopic Partition: 0    Leader: 0   Replicas: 0,1   Isr: 0
    Topic: partopic Partition: 1    Leader: 0   Replicas: 1,0   Isr: 0
```
