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

## gin-docker安装

docker build -t gin-docker . //打包 修改mysql的ip

docker run --link mysql:mysql -p 8001:8001 gin-docker   //运行在8001端口

接口访问127.0.0.1:8001

docker ps -a    //查看本地镜像

docker exec -it 镜像id /bin/bash //进入容器内部
