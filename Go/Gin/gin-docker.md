# 记录下本地使用docker运行gin框架项目


## mysql安装

docker pull mysql:latest  //拉取镜像

docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql  //运行

docker ps -a     //查看本地镜像

docker start 镜像id    //启动

连接Mysql数据库

1.通过工具

2.docker下命令行连接

docker exec -it 镜像id /bin/bash //进入容器内部

mysql -uroot -p -h localhost

## redis安装

docker pull redis:latest    //拉取镜像

docker run -itd --name redis -p 6379:6379 redis  //运行

docker ps -a     //查看本地镜像

docker exec -it 镜像id /bin/bash   //进入容器内部

redis-cli -h 127.0.0.1 -p 6379     //连接redis

## clickhouse安装

1、拉取clickhouse的docker镜像

docker pull yandex/clickhouse-server

docker pull yandex/clickhouse-clinet

2、启动server端

- 默认直接启动即可
docker run -d --name clickhouse-server --ulimit nofile=262144:262144 yandex/clickhouse-server

- 如果想指定目录启动，这里以clickhouse-test-server命令为例，可以随意写

mkdir /work/clickhouse/clickhouse-test-db       ## 创建数据文件目录

- 使用以下路径启动，在外只能访问clickhouse提供的默认9000端口，只能通过clickhouse-client连接server

docker run -d --name clickhouse-server --ulimit nofile=262144:262144 --volume=/work/clickhouse/clickhouse_test_db:/var/lib/clickhouse yandex/clickhouse-server

3、本地连接

docker run -it --rm --link clickhouse-server:clickhouse-server yandex/clickhouse-client --host clickhouse-server

docker run -it -p 9000:9000 yandex/clickhouse-server //以本地9000端口启动

docker exec -it 镜像id /bin/bash //进入容器内部

## gin-docker安装

docker build -t gin-docker . //打包 修改mysql的ip

docker run --link mysql:mysql -p 8001:8001 gin-docker   //运行在8001端口

接口访问127.0.0.1:8001

docker ps -a    //查看本地镜像

docker exec -it 镜像id /bin/bash //进入容器内部
