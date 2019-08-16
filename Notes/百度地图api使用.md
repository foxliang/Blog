## 1.根据经纬度 获取城市信息

接口：http://api.map.baidu.com/geocoder/v2/?ak=ak&location=31.241857,121.49635&output=json

返回：
```
{
    "status": 0,
    "result": {
        "location": {
            "lng": 121.49634999999994,
            "lat": 31.241856867952256
        },
        "formatted_address": "上海市黄浦区中山东一路12号六楼",
        "business": "南京东路,外滩,北京东路",
        "addressComponent": {
            "country": "中国",
            "country_code": 0,
            "country_code_iso": "CHN",
            "country_code_iso2": "CN",
            "province": "上海市",
            "city": "上海市",
            "city_level": 2,
            "district": "黄浦区",
            "town": "",
            "adcode": "310101",
            "street": "中山东一路",
            "street_number": "12号六楼",
            "direction": "附近",
            "distance": "28"
        },
        "pois": [],
        "roads": [],
        "poiRegions": [
            {
                "direction_desc": "内",
                "name": "上海浦东发展银行(总行)",
                "tag": "金融;银行",
                "uid": "3c0ea7c5eda6129dfbf0cb5e",
                "distance": "0"
            }
        ],
        "sematic_description": "上海浦东发展银行(总行)内,中山东一路十二号大楼附近24米",
        "cityCode": 289
    }
}
```

## 2.根据百度地图-获取城市对应经纬度

网址：http://api.map.baidu.com/lbsapi/getpoint/index.html

![image](
https://github.com/weilyf2017/Blog/blob/master/images/%E7%99%BE%E5%BA%A6%E5%9C%B0%E5%9B%BE-%E8%8E%B7%E5%8F%96%E5%9C%B0%E7%82%B9%E7%BB%8F%E7%BA%AC%E5%BA%A6.png)
