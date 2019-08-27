ELK这里想要区分账号给不同的权限访问 需要打开 x-pack（30天试用）

## 1，在 elasticsearch.yml 中 修改/加入

```
xpack.security.enabled: true
```
然后重启elasticsearch

## 2，设置用户名和密码
```
bin/elasticsearch-setup-passwords interactive
```
这时候可能会提示权限不够，win10中需用管理员身份

## 3，再次执行设置用户名和密码的命令,这里需要为4个用户分别设置密码，elastic, kibana, logstash_system,beats_system
```
E:\ELK\elasticsearch-7.2.0\elasticsearch-7.2.0\bin>elasticsearch-setup-passwords interactive
future versions of Elasticsearch will require Java 11; your Java version from [D:\Program Files\java\jdk1.8.0_171\jre] does not meet this requirement
Initiating the setup of passwords for reserved users elastic,apm_system,kibana,logstash_system,beats_system,remote_monitoring_user.
You will be prompted to enter passwords as the process progresses.
Please confirm that you would like to continue [y/N]y


Enter password for [elastic]:
Reenter password for [elastic]:
Enter password for [apm_system]:
Reenter password for [apm_system]:
Enter password for [kibana]:
Reenter password for [kibana]:
Enter password for [logstash_system]:
Reenter password for [logstash_system]:
Enter password for [beats_system]:
Reenter password for [beats_system]:
Enter password for [remote_monitoring_user]:
Reenter password for [remote_monitoring_user]:
Changed password for user [apm_system]
Changed password for user [kibana]
Changed password for user [logstash_system]
Changed password for user [beats_system]
Changed password for user [remote_monitoring_user]
Changed password for user [elastic]
```
## 4，修改kibana配置文件,config下的kibana.yml,添加如下内容
```
elasticsearch.username: "elastic"
elasticsearch.password: "123456"
```
## 5 ，这个时候再重启kibana 登录就需要输出账号密码了
![image](https://img-blog.csdnimg.cn/201908271118127.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)

## 6，这里还能添加不同的用户和权限，来管理elasticsearch和kibana
![image](https://img-blog.csdnimg.cn/20190827112229991.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)

## 7，我这边配置了一个账号test，角色test，权限只有discover的账户

![image](https://img-blog.csdnimg.cn/20190827112420704.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)
 

好了 我本地的测试已经结束了。

参考：

https://www.elastic.co/cn/subscriptions

![image](https://img-blog.csdnimg.cn/20190827111231533.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)

## 注意：这个功能不是免费的，这个只是本人自学研究，商业使用还是要咨询官方，毕竟只是开源没有免费。
