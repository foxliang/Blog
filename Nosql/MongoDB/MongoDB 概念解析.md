不管我们学习什么数据库都应该学习其中的基础概念，在mongodb中基本的概念是文档、集合、数据库，下面我们挨个介绍。

下表将帮助您更容易理解Mongo中的一些概念：

SQL术语/概念 | MongoDB术语/概念 | 解释/说明
---|---|---
database |database |	数据库
table | collection|		数据库表/集合
row |document|数据库表/集合
column |field |	数据字段/域
index |index |	索引
table joins | |	表连接,MongoDB不支持
primary key |primary key |	主键,MongoDB自动将_id字段设置为主键	


## 数据库
一个mongodb中可以建立多个数据库。

MongoDB的默认数据库为"db"，该数据库存储在data目录中。

MongoDB的单个实例可以容纳多个独立的数据库，每一个都有自己的集合和权限，不同的数据库也放置在不同的文件中。

"show dbs" 命令可以显示所有数据的列表。

## 文档(Document)
文档是一组键值(key-value)对(即 BSON)。MongoDB 的文档不需要设置相同的字段，并且相同的字段不需要相同的数据类型，这与关系型数据库有很大的区别，也是 MongoDB 非常突出的特点。

一个简单的文档例子如下：
```
{"site":"www.runoob.com", "name":"菜鸟教程"}
```

## 集合
集合就是 MongoDB 文档组，类似于 RDBMS （关系数据库管理系统：Relational Database Management System)中的表格。

集合存在于数据库中，集合没有固定的结构，这意味着你在对集合可以插入不同格式和类型的数据，但通常情况下我们插入集合的数据都会有一定的关联性。

比如，我们可以将以下不同数据结构的文档插入到集合中：
```
{"site":"www.baidu.com"}
{"site":"www.google.com","name":"Google"}
{"site":"www.runoob.com","name":"菜鸟教程","num":5}
```
当第一个文档插入时，集合就会被创建。

##### 合法的集合名

- 集合名不能是空字符串""。

- 集合名不能含有\0字符（空字符)，这个字符表示集合名的结尾。

- 集合名不能以"system."开头，这是为系统集合保留的前缀。

- 用户创建的集合名字不能含有保留字符。有些驱动程序的确支持在集合名里面包含，这是因为某些系统生成的集合中包含该字符。除非你要访问这种系统创建的集合，否则千万不要在名字里出现$。　

如下实例：
```
db.col.findOne()
```
