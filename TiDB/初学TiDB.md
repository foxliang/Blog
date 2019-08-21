1、tidb在tikv上层，tidb用go实现，tikv用rust实现

2、tidb的目标是实现google的F1，tikv和google的spanner有关系​

3、sql语句解析成AST（Abstract syntax trees），​使用optimizer转换为查询计划，并对查询计划进行逻辑优化（logic.go）和物理优化，变成Plan tree，再由executor转换成Executor tree

4、F1是google开发的一个可容错可扩展的RDBMS，​基于spanner（重要的底层存储技术），支持sql。

5、spanner致力于跨数据中心的数据复制，也提供数据库功能，​是一个“临时多版本”的数据库。数据存储在一个版本化的关系表里，数据会根据其提交的时间打上时间错。spanner有一个全球时间同步机制，可以在数据提交时给出一个时间戳，保证了外部一致性。

6、启动顺序：PD->TIKV->TIDB。假定有3个node，每个node依次启动PD、TIKV，在node1启动TIDB，然后可以用mysql client连接TIDB。PD(placement driver)负责管理协调TIKV cluster。


知乎：https://www.zhihu.com/question/51131241
