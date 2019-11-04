# 使用 BINARY 关键字使搜索区分大小写

```
select id,name from user where BINARY  name='fox';
```

![image](https://img-blog.csdnimg.cn/20191104100403559.png)


## 不加 BINARY  

```
select id,name from user where name='fox';
```

![image](https://img-blog.csdnimg.cn/20191104100252507.png)

