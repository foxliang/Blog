# 搭建一个api和一个mysql服务

- 能够接收用户传的数据并且存在mysql里面

- 数据需要持久化


### 1.部署

```
kubectl apply -f mysql-config.yaml
kubectl apply -f mysql-pv.yaml
kubectl apply -f mysql-pvc.yaml
kubectl apply -f mysql.yaml
kubectl apply -f go.yaml
```

### 2.查看服务

```
$ kubectl get configmap
NAME               DATA   AGE
mysql-config       1      76m

$ kubectl get pv       
NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM               STORAGECLASS    REASON   AGE
mysql-pv     100Mi      RWO            Retain           Bound       default/mysql-pvc   nfs                      76m

$ kubectl get pvc
NAME        STATUS   VOLUME     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
mysql-pvc   Bound    mysql-pv   100Mi      RWO            nfs            76m

$ kubectl get pods
NAME                                    READY   STATUS    RESTARTS   AGE
mysql-fox-deployment-7676755855-q4ghq   1/1     Running   0          6m59s

$ kubectl get svc 
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
mysql-fox    NodePort    10.110.211.206   <none>        3306:31306/TCP   7m2s

$ minikube service list
|----------------------|---------------------------|---------------------|---------------------------|
|      NAMESPACE       |           NAME            |     TARGET PORT     |            URL            |
|----------------------|---------------------------|---------------------|---------------------------|
| default              | mysql-fox                 | mysql-fox-port/3306 | http://192.168.79.2:31306 |
```
### 3.连接mysql

```
# mysql -h192.168.79.2 -P31306 -uroot -p
123456

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.01 sec)

mysql> create database test;
Query OK, 1 row affected (0.10 sec)

mysql> use test;
Database changed
mysql> show tables;
Empty set (0.00 sec)

mysql> CREATE TABLE `go` (
    ->   `id` int unsigned NOT NULL AUTO_INCREMENT,
    ->   `name` varchar(10) NOT NULL DEFAULT '' COMMENT 'name',
    ->   `create_time` int unsigned NOT NULL DEFAULT '0' COMMENT '',
    ->   `update_time` int unsigned NOT NULL DEFAULT '0' COMMENT '',
    ->   `up_timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    ->   PRIMARY KEY (`id`),
    ->   KEY `idx_create_time` (`create_time`)
    -> ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='test';
Query OK, 0 rows affected (0.64 sec)

mysql> show tables;
+----------------+
| Tables_in_test |
+----------------+
| go             |
+----------------+
1 row in set (0.01 sec)

mysql> select * from go;
Empty set (0.01 sec)

```

### 4.连接api
```
#新增

$ curl  http://192.168.79.2:31088/data/create\?name\=1
"success create"                                      

#获取单条

$ curl  http://192.168.79.2:31088/data/get\?name\=1

{"id":1,"name":"1","create_time":1619507317,"update_time":1619507317}

#获取列表

$ curl  http://192.168.79.2:31088/data/list        
 
[{"id":1,"name":"1","create_time":1619507317,"update_time":1619507317},{"id":2,"name":"33","create_time":1619507325,"update_time":1619507627},{"id":3,"name":"33","create_time":1619507639,"update_time":1619507639},{"id":4,"name":"44","create_time":1619507902,"update_time":1619507902},{"id":5,"name":"44","create_time":1619508242,"update_time":1619508242},{"id":6,"name":"23","create_time":1619510263,"update_time":1619510263},{"id":7,"name":"7","create_time":1619517602,"update_time":1619517602}]       
```


### 5.测试容器,节点停止再重启看数据是否存在


如遇到的问题

1. Host '172.17.0.1' is not allowed to connect to this MySQL server

说明所连接的用户帐号没有远程连接的权限，只能在本机(localhost)登录。 需更改 mysql 数据库里的 user表里的 host项

把localhost改称%

```
进入mysql

mysql> use mysql

mysql> update user set host='%' where user = 'root';

mysql> select host from user where user = 'root';

+-----------------------+

|host|

+-----------------------+

|% |

MySQL> flush privileges;

```

2. Access denied for user 'root'@'172.17.0.1' (using password: YES)

说明密码不对,重新设置或者填写正确的密码,默认为空
