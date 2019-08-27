ELK这里想要区分账号给不同的权限访问 需要打开 x-pack（30天试用）

## 1，在 elasticsearch.yml 中 修改/加入

xpack.security.enabled: true

然后重启elasticsearch

## 2，设置用户名和密码
bin/elasticsearch-setup-passwords interactive
这时候可能会提示权限不够，win10中需用管理员身份

## 3，再次执行设置用户名和密码的命令,这里需要为4个用户分别设置密码，elastic, kibana, logstash_system,beats_system


## 4，修改kibana配置文件,config下的kibana.yml,添加如下内容
elasticsearch.username: "elastic"
elasticsearch.password: "123456"
## 5 ，这个时候再重启kibana 登录就需要输出账号密码了


## 6，这里还能添加不同的用户和权限，来管理elasticsearch和kibana


## 7，我这边配置了一个账号test，角色test，权限只有discover的账户


 

好了 我本地的测试已经结束了。

参考：

https://www.elastic.co/cn/subscriptions



## 注意：这个功能不是免费的，这个只是本人自学研究，商业使用还是要咨询官方，毕竟只是开源没有免费。
