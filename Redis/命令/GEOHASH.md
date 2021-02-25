Redis geohash 用于返回一个或多个位置元素的 Geohash 表示。Redis GEO 使用 geohash 来保存地理位置的坐标。

语法

geohash 语法格式如下：

```
GEOHASH key member [member ...]
```

命令:

```
127.0.0.1:6379> geoadd company 116.48105 39.996794 fox
(integer) 1
127.0.0.1:6379> geoadd company 114.48105 39.996794 lina
(integer) 1
127.0.0.1:6379> geoadd company 114.48105 36.996794 mask
(integer) 1
127.0.0.1:6379> geoadd company 114.44105 36.946794 bili
(integer) 1
127.0.0.1:6379> geoadd company 114.14105 36.646794 jack 120.50305 40.864512 tony
(integer) 2
127.0.0.1:6379> geodist company fox mask km
"376.3396"
127.0.0.1:6379> geodist company fox mask m
"376339.5580"
127.0.0.1:6379> geodist company mask tony km
"675.3889"
127.0.0.1:6379> geopos company fox
1) 1) "116.48104995489120483"
   2) "39.99679348858259686"
127.0.0.1:6379> geopos company tony
1) 1) "120.50304919481277466"
   2) "40.86451218722258005"
127.0.0.1:6379> georadiusbymember company fox 20km count 3 asc
(error) ERR need numeric radius
127.0.0.1:6379> georadiusbymember company fox 20 km count 3 asc
1) "fox"
127.0.0.1:6379> georadiusbymember company fox 200 km count 3 asc
1) "fox"
2) "lina"
127.0.0.1:6379> georadiusbymember company fox 2000 km count 3 asc
1) "fox"
2) "lina"
3) "tony"
127.0.0.1:6379> georadiusbymember company fox 20000 km count 3 asc
1) "fox"
2) "lina"
3) "tony"
127.0.0.1:6379> georadiusbymember company fox 20000 km count 3 desc
1) "jack"
2) "bili"
3) "mask"
```

- 增加

geoadd 指令携带集合名称以及多个经纬度名称三元组,注意这里可以加入多个三元组

- 距离

geodist 指令可以用来计算两个元素之间的距离,携带集合名称、2 个名称和距离单位。

- 获取元素位置

geopos 指令可以获取集合中任意元素的经纬度坐标,可以一次获取多个。

