# Redis
下载地址：https://download.csdn.net/download/qq_35349114/12101406

也可以直接到github下载 ：https://github.com/MicrosoftArchive/redis/releases

下载稳定版本，然后解压到本地，在文件夹中点击 redis-server.exe 就运行Redis服务端了

![image](https://github.com/foxliang/Blog/blob/master/images/win10%E5%AE%89%E8%A3%85Redis%E5%92%8C%E5%8F%AF%E8%A7%86%E5%8C%96%E5%AE%A2%E6%88%B7%E7%AB%AF/1.png)

![image](https://github.com/foxliang/Blog/blob/master/images/win10%E5%AE%89%E8%A3%85Redis%E5%92%8C%E5%8F%AF%E8%A7%86%E5%8C%96%E5%AE%A2%E6%88%B7%E7%AB%AF/2.png)

![image](https://github.com/foxliang/Blog/blob/master/images/win10%E5%AE%89%E8%A3%85Redis%E5%92%8C%E5%8F%AF%E8%A7%86%E5%8C%96%E5%AE%A2%E6%88%B7%E7%AB%AF/3.png)

在这里就把redis服务启动了，但是蛋疼的是关闭窗口之后服务就消失了，原因在于这个服务没有配置到win10下面，也可以配置到服务中这样每次启动就打开了

#测试 在当前目录

运行命令 启动redis客户端

```
 redis-cli.exe -h 127.0.0.1 -p 6379
```


 ![image](https://github.com/foxliang/Blog/blob/master/images/win10%E5%AE%89%E8%A3%85Redis%E5%92%8C%E5%8F%AF%E8%A7%86%E5%8C%96%E5%AE%A2%E6%88%B7%E7%AB%AF/4.png)

## 可视化工具

github地址：https://github.com/caoxinyu/RedisClient

连接到本地ip：127.0.0.1


![image](https://github.com/foxliang/Blog/blob/master/images/win10%E5%AE%89%E8%A3%85Redis%E5%92%8C%E5%8F%AF%E8%A7%86%E5%8C%96%E5%AE%A2%E6%88%B7%E7%AB%AF/5.png)

查看对应数据库的数据

![image](https://github.com/foxliang/Blog/blob/master/images/win10%E5%AE%89%E8%A3%85Redis%E5%92%8C%E5%8F%AF%E8%A7%86%E5%8C%96%E5%AE%A2%E6%88%B7%E7%AB%AF/6.png)

 
