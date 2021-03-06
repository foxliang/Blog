# ELK 是 elastic 公司旗下三款产品 ElasticSearch 、Logstash 、Kibana 的首字母组合。
Elasticsearch简称ES，是一个基于Lucene的实时分布式的搜索与分析引擎，是遵从Apache开源条款的一款开源产品，是当前主流的企业级搜索引擎。它提供了一个分布式服务，可以使您快速的近乎于准实时的存储、查询和分析超大数据集，通常被用来当做构建复杂查询特性和需求强大应用的基础引擎或技术。

### ElasticSearch 可以被用在如下几个场景：
1. 当你运营一个提供客户检索商品的在线电子商城的时候，可以使用ES来存储整个商品目录和库存，并且为客户提供检索和自动推荐功能。

2. 收集交易数据，存储并做趋势、统计、概要或异常分析。这种情况下，可以使用Logstash来收集、聚合和解析数据，并且存储到 Elasticsearch。一单数据进入 Elasticsearch，你可以检索，聚合来掌握你感兴趣的信息。 

3. 价格预警平台，为价格敏感客户提供匹配其需求（主要是价格方面）的商品。 

4. 在报表分析/BI领域，可以使用ES的聚合功能完成针对大数据量的复杂分析。

#### 1、ElasticSearch是个开源分布式搜索引擎，提供搜集、分析、存储数据三大功能。

它的特点有：分布式，零配置，自动发现，索引自动分片，索引副本机制，restful风格接口，多数据源，自动搜索负载等。

#### 2、Logstash 主要是用来日志的搜集、分析、过滤日志的工具，支持大量的数据获取方式。

一般工作方式为c/s架构，client端安装在需要收集日志的主机上，server端负责将收到的各节点日志进行过滤、修改等操作在一并发往elasticsearch上去。

#### 3、Kibana可以为 Logstash 和 ElasticSearch 提供的日志分析友好的 Web 界面，可以帮助汇总、分析和搜索重要数据日志。


官网：https://www.elastic.co/cn/

下载地址：https://www.elastic.co/cn/downloads/

中文社区：https://elasticsearch.cn/

ElasticSearch: 权威指南：https://www.elastic.co/guide/cn/elasticsearch/guide/current/index.html

ElasticSearch-PHP：https://www.elastic.co/guide/cn/elasticsearch/php/current/index.html

Kibana 用户手册：https://www.elastic.co/guide/cn/kibana/current/index.html
