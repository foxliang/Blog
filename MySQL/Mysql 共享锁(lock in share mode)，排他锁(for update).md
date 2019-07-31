## 共享锁(lock in share mode)
#### 简介
允许不同事务之前共享加锁读取，但不允许其它事务修改或者加入排他锁
如果有修改必须等待一个事务提交完成，才可以执行，容易出现死锁

#### 共享锁事务之间的读取
session1:

```
start transaction;
select * from test where id = 1 lock in share mode;
```
结果：

![image](https://img-blog.csdnimg.cn/20190731140214666.png)

session2:

```
start transaction;
select * from test where id = 1 lock in share mode;
```
结果：

![image](https://img-blog.csdnimg.cn/20190731140214666.png)

<!--more-->

此时session1和session2都可以正常获取结果，那么再加入session3 排他锁读取尝试

session3:

```
start transaction;
select * from test where id = 1 for update;
```
结果：

![image](https://img-blog.csdnimg.cn/20190731140303837.png)


在session3中则无法获取数据，直到超时或其它事物commit

```
Lock wait timeout exceeded; try restarting transaction
```

#### 共享锁之间的更新
当session1执行了修改语句：
session1:

```
update test set name = 'php7' where id = 1;
```
可以很多获取执行结果。
当session2再次执行修改id=1的语句时：
session2:

```
update test set name = 'mysql8' where id = 1;
```
就会出现死锁或者锁超时，错误如下：

```
Deadlock found when trying to get lock; try restarting transaction
```
或者：

```
Lock wait timeout exceeded; try restarting transaction
```
必须等到session1完成commit动作后，session2才会正常执行，如果此时多个session并发执行，可想而知出现死锁的几率将会大增。

session3则更不可能

结论：
mysql共享锁(
lock in share mode
)
允许其它事务也增加共享锁读取
不允许其它事物增加排他锁(
for update
)
当事务同时增加共享锁时候，事务的更新必须等待先执行的事务commit后才行，如果同时并发太大可能很容易造成死锁
共享锁，事务都加，都能读。修改是惟一的，必须等待前一个事务commit，才可以

## 排他锁(for update)
#### 简介
当一个事物加入排他锁后，不允许其他事务加共享锁或者排它锁读取，更加不允许其他事务修改加锁的行。

#### 排他锁不同事务之间的读取
同样以不同的session来举例
session1:

```
start transaction;
select * from test where id = 1 for update;
```
结果：

![image](https://img-blog.csdnimg.cn/20190731140214666.png)

session2:

```
start transaction;
select * from test where id = 1 for update;
```
结果：

![image](https://img-blog.csdnimg.cn/20190731140754173.png)

当session1执行完成后，再次执行session2，此时session2也会卡住，无法立刻获取查询的数据。直到出现超时

```
Lock wait timeout exceeded; try restarting transaction
```
或session1 commit才会执行

那么再使用session3 加入共享锁尝试

```
select * from test where id = 1 lock in share mode;
```
结果也是如此，和session2一样，超时或等待session1 commit

```
Lock wait timeout exceeded; try restarting transaction
```
#### 排他锁事务之间的修改
当在session1中执行update语句：

```
update test set name = 123 where id = 1;
```
可以正常获取结果

```
Query OK, 1 row affected (0.00 sec)
Rows matched: 1  Changed: 1  Warnings: 0
```
此时在session2中执行修改

```
update test set name = '456' where id = 1;
```
则会卡住直接超时或session1 commit,才会正常吐出结果

session3也很明显和session2一样的结果，这里就不多赘述

### 总结
1.事务之间不允许其它排他锁或共享锁读取，修改更不可能

2.一次只能有一个排他锁执行commit之后，其它事务才可执行

3.不允许其它事务增加共享或排他锁读取。修改是惟一的，必须等待前一个事务commit，才可以
