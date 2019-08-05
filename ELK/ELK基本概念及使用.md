引言：
对于刚接触ES的童鞋，经常搞不明白ES的各个概念的含义。尤其对“索引”二字更是与关系型数据库混淆的不行。本文通过对比关系型数据库，将ES中常见的增、删、改、查操作进行图文呈现。能加深你对ES的理解。同时，也列举了kibana下的图形化展示。

## ES Restful API GET、POST、PUT、DELETE、HEAD含义： 

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

### ////批量数据处理 添加 删除
bulk 与其他的请求体格式稍有不同，如下所示：
```
{ action: { metadata }}\n
{ request body        }\n
{ action: { metadata }}\n
{ request body        }\n
'
```
 这种格式类似一个有效的单行 JSON 文档 流 ，它通过换行符(\n)连接到一起。注意两个要点：

每行一定要以换行符(\n)结尾， 包括最后一行 。这些换行符被用作一个标记，可以有效分隔行。

这些行不能包含未转义的换行符，因为他们将会对解析造成干扰。这意味着这个 JSON 不 能使用 pretty 参数打印。

**action/metadata** 行指定 哪一个文档 做 什么操作 。

action 必须是以下选项之一:

**create**

如果文档不存在，那么就创建它。详情请见 创建新文档。

**index**

创建一个新文档或者替换一个现有的文档。详情请见 索引文档 和 更新整个文档。

**update**

部分更新一个文档。详情请见 文档的部分更新。

**delete**

删除一个文档。详情请见 删除文档。

**metadata** 应该 指定被索引、创建、更新或者删除的文档的 _index 、 _type 和 _id 

1.指定index
```
curl -XPOST "http://localhost:9200/zhuita-test-2019-08-02/_doc/_bulk" -H 'Content-Type: application/json' -d'
{"name" : "fox4","date":1564724394000 }
{"name" : "fox5","date":1564724494000 }
{"name" : "fox6","date":1564724594000 }
'
```
2.指定通用模式 
```curl -XPOST "/_bulk" -H 'Content-Type: application/json' -d'
{ "index": { "_index": "zhuita-test-2019-08-02", "_type": "_doc" }}
{ "name" : "fox1","date":1564724294000 } 
{ "index": { "_index": "zhuita-test-2019-08-02", "_type": "_doc" }}
{ "name" : "fox2","date":1564724394000 } 
'
```
3.为了把所有的操作组合在一起，一个完整的 bulk 请求 有以下形式:
```curl -XPOST "/_bulk" -H 'Content-Type: application/json' -d'
{ "delete": { "_index": "website", "_type": "blog", "_id": "123" }} 
{ "create": { "_index": "website", "_type": "blog", "_id": "123" }}
{ "title":    "My first blog post" }
{ "index":  { "_index": "website", "_type": "blog" }}
{ "title":    "My second blog post" }
{ "update": { "_index": "website", "_type": "blog", "_id": "123", "_retry_on_conflict" : 3} }
{ "doc" : {"title" : "My updated blog post"} } 
'
```

### 在kibana查看分析
![image](https://img-blog.csdnimg.cn/20190730102745977.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1MzQ5MTE0,size_16,color_FFFFFF,t_70)
