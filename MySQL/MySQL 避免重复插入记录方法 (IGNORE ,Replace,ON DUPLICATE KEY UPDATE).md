MySQL唯一键 unique key，用来保证对应的字段中的数据唯一的。

主键也可以用保证字段数据唯一性，但是一张表只有一个主键。

唯一键特点：

1、唯一键在一张表中可以有多个。

2、唯一键允许字段数据为NULL，NULL可以有多个（NULL不参与比较）

有时候会遇到主键和唯一键的冲突，这时候需要下面几种方法

### 1.INSERT IGNORE INTO 
当插入数据时，如出现错误时，如重复数据，将不返回错误，只以警告形式返回。所以使用ignore请确保语句本身没有问题，否则也会被忽略掉

返回行数0

### 2.ON DUPLICATE KEY UPDATE
当primary或者unique重复时，则执行update语句，如update后为无用语句，如id=id，则同1功能相同，但错误不会被忽略掉

返回行数0

### 3.REPLACE INTO
如果存在primary or unique相同的记录，则先删除掉。再插入新记录。

返回行数2
