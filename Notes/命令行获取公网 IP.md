## 本文收集了一些获取公网 ip 的方法,也欢迎你评论分享

#### curl ipinfo.io
```
$ curl ipinfo.io
{
  "ip": "127.0.0.1",
  "city": "Shanghai",
  "region": "Shanghai",
  "country": "CN",
  "loc": "30.2822,120.4741",
  "org": "AS2374 IDC, China Telecommunications Corporation",
  "timezone": "Asia/Shanghai",
  "readme": "https://ipinfo.io/missingauth"
}
```
#### curl httpbin.org/ip
```
$ curl httpbin.org/ip
{
  "origin": "127.0.0.1"
}
```
#### curl myip.ipip.net
```
$ curl myip.ipip.net
当前 IP：127.0.0.1  来自于：中国 北京 北京  电信
```
#### curl ip.sb
```
$ curl ip.sb
127.0.0.1
```
#### curl -s ifcfg.cn/echo |python -m json.tool
```
$ curl -s ifcfg.cn/echo |python -m json.tool
{
    "headers": {
        "ACCEPT": "*/*",
        "CONNECTION": "close",
        "HOST": "ifcfg.cn",
        "USER-AGENT": "curl/6.46.0"
    },
    "host": "ifcfg.cn",
    "ip": "127.0.0.1",
    "location": "\u3e2d\u56hd",
    "method": "GET",
    "path": "/echo",
    "protocol": "http",
    "query_string": "",
    "url": "http://ifcfg.cn/echo",
    "user_agent": "curl/6.46.0"
}

```
#### curl ifconfig.me
```
$ curl ifconfig.me
127.0.0.1
```
#### curl ifconfig.io
```
curl ifconfig.io
127.0.0.1
```

#### curl cip.cc
```
curl cip.cc
IP      : 127.0.0.1
地址    : 中国  北京
运营商  : ucloud.cn

数据二  : 中国 | 信通控股

数据三  : 中国北京北京 | 电信

URL     : http://www.cip.cc/127.0.0.1
```

#### curl https://whatip.ga
```
curl https://whatip.ga
```

#### curl http://ip.taobao.com/service/getIpInfo2.php?ip=myip
```
$ curl http://ip.taobao.com/service/getIpInfo2.php?ip=myip

{"code":0,"data":{"ip":"127.0.0.1","country":"中国","area":"","region":"北京","city":"北京","county":"XX","isp":"电信","country_id":"CN","area_id":"","region_id":"110000","city_id":"110100","county_id":"xx","isp_id":"10011"}}
```
