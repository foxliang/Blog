## 删除 null image

sudo docker rmi $(docker images -f "dangling=true" -q)

删除所有镜像

## 删掉容器

docker stop $(docker ps -qa)
docker rm $(docker ps -qa)

## 删除镜像

docker rmi --force $(docker images -q)

删除名称中包含某个字符串的镜像

## 例如删除包含“some”的镜像

docker rmi --force $(docker images | grep some | awk '{print $3}')
