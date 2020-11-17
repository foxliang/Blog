
这里直接选用 sebp/elk

## 1.拉镜像
```
docker pull sebp/elk
```
## 2.运行并进入容器

```
docker run -p 5601:5601 -p 9200:9200 -p 5044:5044 -e ES_MIN_MEM=128m -e ES_MAX_MEM=1024m -it --name elk sebp/elk
```

## 3.浏览器查看 http://127.0.0.1:9200/

```
{
  "name" : "elk",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "05PWNSBDRJSedkPvavDlrg",
  "version" : {
    "number" : "7.9.3",
    "build_flavor" : "oss",
    "build_type" : "tar",
    "build_hash" : "c4138e51121ef06a6404866cddc601906fe5c868",
    "build_date" : "2020-10-16T10:36:16.141335Z",
    "build_snapshot" : false,
    "lucene_version" : "8.6.2",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```


4.运行代码创建记录

代码文件： [https://github.com/foxliang/Blog/blob/master/Go/my_go/elk/main.go](https://github.com/foxliang/Blog/blob/master/Go/my_go/elk/main.go)

![images](https://github.com/foxliang/Blog/blob/master/images/elk_demo.png)

