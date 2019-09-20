# 使用GO 爬虫--爬取豆瓣电影top

网址：https://movie.douban.com/top250

1.[存到csv文件中](https://github.com/foxliang/Blog/blob/master/Go/my_go/douban/douban.go)

2.[存到mysql中](https://github.com/foxliang/Blog/blob/master/Go/my_go/douban/test2.go)

表结构：
```
CREATE TABLE `douban` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '电影名字',
  `average` decimal(20,2) DEFAULT NULL COMMENT '评分',
  `commentCount` bigint(255) DEFAULT NULL COMMENT '评价人数',
  `careted_at` datetime DEFAULT NULL COMMENT '时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```
