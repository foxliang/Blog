## ElasticSearch采用REST API，所有的操作都可通过HTTP API完成，例如增删改查、别名配置等。
[RESTful API 设计的总结](https://github.com/weilyf2017/Blog/blob/master/Notes/RESTful%20API%20%E8%AE%BE%E8%AE%A1%E7%9A%84%E6%80%BB%E7%BB%93.md)

```math
Elasticsearch是一个基于Apache Lucene(TM)的开源搜索引擎。

无论在开源还是专有领域， Lucene可以被认为是迄今为止最先进、性能最好的、功能最全的搜索引擎库。 

但是，Lucene只是一个库。想要使用它，你必须使用Java来作为开发语言并将其直接集成到 你的应用中，更糟糕的是，Lucene非常复杂，你需要深入了解检索的相关知识来理解它是如 何工作的。Elasticsearch也使用Java开发并使用Lucene作为其核心来实现所有索引和搜索的功能，
但是 它的目的是通过简单的 RESTful API 来隐藏Lucene的复杂性，从而让全文搜索变得简。
```

### 使用REST Client交互
客户端访问仅支持HTTP / TCP方式，建议您采用Elasticsearch官方提供的Java REST Client。

### 使用Java API交互
Elasticsearch为Java用户提供了内置客户端。关于Java API的更多信息，请查看官方Java API文档。

### 传输客户端（Transport client）

传输客户端能够发送请求到远程集群，它自己不加入集群，只是简单转发请求给集群中的节点。

传输客户端通过9300端口与集群交互，使用Elasticsearch传输协议（Elasticsearch Transport Protocol）。

集群中的节点之间也通过9300端口进行通信。如果此端口未开放，您的节点将不能组成集群。

说明 Java客户端所在的Elasticsearch版本必须与集群中其他节点一致，否则它们可能无法相互识别。
### RESTful API（HTTP）
其他所有程序语言都可以使用RESTful API，通过9200端口与Elasticsearch进行通信。可使用您喜欢的Web客户端，或通过curl命令与Elasticsearch通信。

说明
Elasticsearch官方提供了多种程序语言的客户端，例如Groovy、Javascript、.NET、PHP、Perl、Python以及Ruby。

还有很多由社区提供的客户端和插件，您可以在官方文档中获取。

curl请求组成（HTTP）

```
curl -X<VERB> '<PROTOCOL>://<HOST>:<PORT>/<PATH>?<QUERY_STRING>' -d '<BODY>'
```
VERB：HTTP方法，包括GET、POST、PUT、HEAD、DELETE。

PROTOCOL：http或者https协议（只有在Elasticsearch前面有https代理的时候可用）。

HOST：Elasticsearch集群中的任何一个节点的主机名，如果是在本地的节点，那么可使用localhost。

PORT：Elasticsearch HTTP服务所在的端口，默认为9200。

PATH：API路径（例如_count将返回集群中文档的数量），PATH可以包含多个组件，例如_cluster/stats或者_nodes/stats/jvm。

QUERY_STRING： 一些可选的查询请求参数，例如?pretty参数可使请求返回的JSON数据更加美观易读。

BODY：一个JSON格式的请求主体（如果请求需要的话）。

示例

#### 1.统计`Elasticserach`集群中文档数命令：

```
curl -XGET http://localhost:9200/_count?pretty
```

返回：

```
{
  "count" : 86547,
  "_shards" : {
    "total" : 5,
    "successful" : 5,
    "skipped" : 0,
    "failed" : 0
  }
}
```

#### 2.创建一个index为`user`的索引，name为`fox`的数据

```
curl -X POST \
http://10.9.183.17:9200/user/info \
-H 'Content-Type: application/json' \
-d '{
    "name" : "fox"
}'
```
返回：

```
{"_index":"user","_type":"info","_id":"kPjEKGwBjYq9m-uKR-ty","_version":1,"result":"created","_shards":{"total":2,"successful":2,"failed":0},"_seq_no":2,"_primary_term":1}
```

#### 3.修改`user`的索引，name为`Fox`的数据

```
curl -X PUT \
http://10.9.183.17:9200/user/info/kPjEKGwBjYq9m-uKR-ty \
-H 'Content-Type: application/json' \
-d '{
    "name" : "Fox"
}'
```
返回：

```
{"_index":"user","_type":"info","_id":"kPjEKGwBjYq9m-uKR-ty","_version":2,"result":"updated","_shards":{"total":2,"successful":2,"failed":0},"_seq_no":4,"_primary_term":1}
```

#### 4.查询`user`的索引，url之后加`pretty` 是为了美化返回结果

```
 curl -s -XGET http://10.9.183.17:9200/user/_search?pretty
```
返回：

```
{
  "took" : 1,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 4,
      "relation" : "eq"
    },
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "user",
        "_type" : "info",
        "_id" : "evjCKGwBjYq9m-uKxevJ",
        "_score" : 1.0,
        "_source" : {
          "query" : {
            "match_all" : { }
          }
        }
      },
      {
        "_index" : "user",
        "_type" : "info",
        "_id" : "f_jDKGwBjYq9m-uKGuud",
        "_score" : 1.0,
        "_source" : {
          "name" : "fox"
        }
      },
      {
        "_index" : "user",
        "_type" : "info",
        "_id" : "Z_jFKGwBjYq9m-uK5exQ",
        "_score" : 1.0,
        "_source" : {
          "name" : "fox2"
        }
      },
      {
        "_index" : "user",
        "_type" : "info",
        "_id" : "kPjEKGwBjYq9m-uKR-ty",
        "_score" : 1.0,
        "_source" : {
          "name" : "Fox"
        }
      }
    ]
  }
}
```
