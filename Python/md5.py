## 签名计算方法

1.将所有请求参数(key，value 为一组)，对数据结构按照 key 的升序，重新排序, 需要对 null 值进行过滤。

2.将生成好的数据加上秘钥，形成新的数据。

3.对数据使用http_build_query，urldecode处理。

4.生成的字符串 str 做 md5 加密，得到32位大写的sign。

```javascript
import hashlib # 导入模块hashlib应用md5模块
import urllib # urllib提供了一系列用于操作URL的功能
import sys # 系统特定的参数和功能


def md5(str):
    # 创建md5对象
    md5 = hashlib.md5()
    md5.update(str.encode('utf-8'))  # 传入需要加密的字符串进行MD5加密
    return md5.hexdigest()  # 获取到经过MD5加密的字符串并返回

def sign(data):
    data = dict(sorted(data.items(), key=lambda e: e[0])) # 根据键值排序
    # 需要对 null 值进行过滤
    for e in list(data.keys()):
        if data[e] == '' or data[e] == 'null' or data[e] == 'undefined':
            del data[e]
    data['key'] = '2B4C2C023A00E7C6E7D68DA4F650F424' # 秘钥
    sign= urllib.parse.urlencode(data) # urlencode处理 app_id=03CEC941-2E22-7E0D-C7BF-6B1DB772BBA3&ip=127.0.0.1&name=fox&time=1560755194&uid=100&key=2B4C2C023A00E7C6E7D68DA4F650F424
    sign = md5(sign) # md5加密
    return sign.upper() # 把所有字符中的小写字母转换成大写字母

def main():
    # 你要加密的东西
    data = {
        'uid': 100,
        'name': 'fox',
        'ip': '127.0.0.1',
        'time': '1560755194',
        'app_id': '03CEC941-2E22-7E0D-C7BF-6B1DB772BBA3',
        'sign': ''
    }
    return sign(data)


if __name__ == '__main__':

    sign = main() # 执行主函数
    print(sign) # 0F0F9CB7A559968FE161EA2289DF14C5


```
