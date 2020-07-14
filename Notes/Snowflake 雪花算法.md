# Snowflake
![image](https://github.com/foxliang/Blog/blob/master/images/snowflake.png)
算法描述：指定机器 & 同一时刻 & 某一并发序列，是唯一的。据此可生成一个64 bits的唯一ID（long）。默认采用上图字节分配方式：

### sign(1bit)
固定1bit符号标识，即生成的UID为正数。

### delta seconds (28 bits)
当前时间，相对于时间基点"2016-05-20"的增量值，单位：秒，最多可支持约8.7年

### worker id (22 bits)
机器id，最多可支持约420w次机器启动。内置实现为在启动时由数据库分配，默认分配策略为用后即弃，后续可提供复用策略。

### sequence (13 bits)
每秒下的并发序列，13 bits可支持每秒8192个并发。

下面们用golang来实现一下snowFlake

