## 本文收集了一些获取公网 ip 的方法,也欢迎你评论分享

#### curl ipinfo.io
```
$ curl ipinfo.io
{
  "ip": "36.10.25.4",
  "city": "Hangzhou",
  "region": "Zhejiang",
  "country": "CN",
  "loc": "30.2936,120.1614",
  "org": "AS4134 CHINANET-BACKBONE"
}
```
#### curl httpbin.org/ip
```
$ curl httpbin.org/ip
{
  "origin": "36.10.25.4"
}
```
#### curl myip.ipip.net
```
$ curl myip.ipip.net
当前 IP：36.10.25.4  来自于：中国 北京 北京  鹏博士
```
#### curl ip.sb
```
$ curl ip.sb
36.10.25.4
```
#### curl -s ifcfg.cn/echo |python -m json.tool
```
$ curl -s ifcfg.cn/echo |python -m json.tool

{
    "url": "http://ifcfg.cn/echo",
    "user_agent": "curl/7.30.0",
    "protocol": "http",
    "query_string": "",
    "ip": "36.10.25.44",
    "headers": {
        "CONNECTION": "close",
        "HOST": "ifcfg.cn",
        "ACCEPT": "*/*",
        "USER-AGENT": "curl/7.30.0"
    },
    "location": "\u4e2d\u56fd \u5317\u4eac",
    "method": "GET",
    "path": "/echo",
    "host": "ifcfg.cn"
}
```
#### curl ifconfig.me
```
$ curl ifconfig.me
36.10.25.4
```
#### curl ifconfig.io
```
curl ifconfig.io
127.0.0.1
```

#### curl https://whatip.ga
```
curl https://whatip.ga
```

#### curl http://ip.taobao.com/service/getIpInfo2.php?ip=myip
```
$ curl -s http://ip.taobao.com/service/getIpInfo2.php?ip=myip|python -m json.to
ol
{
    "code": 0,
    "data": {
        "country": "\u4e2d\u56fd",
        "country_id": "CN",
        "area": "\u534e\u5317",
        "area_id": "100000",
        "region": "\u5317\u4eac\u5e02",
        "region_id": "110000",
        "city": "\u5317\u4eac\u5e02",
        "city_id": "110100",
        "county": "",
        "county_id": "-1",
        "isp": "\u9e4f\u535a\u58eb",
        "isp_id": "1000143",
        "ip": "36.10.25.44"
    }
}
```
