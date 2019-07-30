引言：
对于刚接触ES的童鞋，经常搞不明白ES的各个概念的含义。尤其对“索引”二字更是与关系型数据库混淆的不行。本文通过对比关系型数据库，将ES中常见的增、删、改、查操作进行图文呈现。能加深你对ES的理解。同时，也列举了kibana下的图形化展示。

ES Restful API GET、POST、PUT、DELETE、HEAD含义： 

1）GET：获取请求对象的当前状态。 

2）POST：改变对象的当前状态。 

3）PUT：创建一个对象。 

4）DELETE：销毁对象。 

5）HEAD：请求获取对象的基础信息。

## Mysql与Elasticsearch核心概念对比示意图

![image](https://img-blog.csdnimg.cn/20190730100818355.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)

### ////新建一个索引，定义date为日期格式来做排序字段
 
```
curl -XPUT "http://localhost:9200/test" -H 'Content-Type: application/json' -d'
{
  "mappings": {
      "properties": {
        "date": {
          "type":   "date",
          "format": "strict_date_optional_time||epoch_millis"
        }
      }
  }
}'
```

### ////创建一个文档 并加入date时间 目前使用毫秒级时间戳
 
```
curl -XPOST "http://localhost:9200/test/_doc" -H 'Content-Type: application/json' -d'
{
  "name":"fox",
  "date":1564137540000
}'
```

### 在kibana查看分析
![image](https://img-blog.csdnimg.cn/20190730102745977.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)
