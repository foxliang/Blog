# WebSocket 简介


Websocket 定义 参考规范 rfc6455


##### 规范解释
```
Websocket 是一种提供客户端(提供不可靠秘钥)与服务端(校验通过该秘钥)进行双向通信的协议。
```
在没有websocket协议之前，要提供客户端与服务端实时双向推送消息，就会使用polling技术，客户端通过xhr或jsonp 发送消息给服务端，并通过事件回调来接收服务端消息。

这种技术虽然也能保证双向通信，但是有一个不可避免的问题，就是性能问题。客户端不断向服务端发送请求，如果客户端并发数过大，无疑导致服务端压力剧增。因此，websocket就是解决这一痛点而诞生的。

这里再延伸一些名词:

- 长轮询

客户端向服务端发送xhr请求,服务端接收并hold该请求，直到有新消息push到客户端，才会主动断开该连接。然后，客户端处理该response后再向服务端发起新的请求。以此类推。

```
HTTP1.1默认使用长连接，使用长连接的HTTP协议，会在响应头中加入下面这行信息: Connection:keep-alive
```

- 短轮询:

客户端不管是否收到服务端的response数据，都会定时想服务端发送请求，查询是否有数据更新。

- 长连接

指在一个TCP连接上可以发送多个数据包，在TCP连接保持期间，如果没有数据包发送，则双方就需要发送心跳包来维持此连接。


```
连接过程: 建立连接——数据传输——...——(发送心跳包，维持连接)——...——数据传输——关闭连接
```


- 短连接

指通信双方有数据交互时，建立一个TCP连接，数据发送完成之后，则断开此连接。


```
连接过程: 建立连接——数据传输——断开连接——...——建立连接——数据传输——断开连接
```

Tips

```
这里有一个误解，长连接和短连接的概念本质上指的是传输层的TCP连接，因为HTTP1.1协议以后，连接默认都是长连接。没有短连接说法(HTTP1.0默认使用短连接)，网上大多数指的长短连接实质上说的就是TCP连接。

http使用长连接的好处: 当我们请求一个网页资源的时候，会带有很多js、css等资源文件，如果使用的时短连接的话，就会打开很多tcp连接，如果客户端请求数过大，导致tcp连接数量过多，对服务端造成压力也就可想而知了。
```


- 单工

数据传输的方向唯一，只能由发送方向接收方的单一固定方向传输数据。

- 半双工

即通信双方既是接收方也是发送方，不过，在某一时刻只能允许向一个方向传输数据。

- 全双工:

即通信双方既是接收方也是发送方，两端设备可以同时发送和接收数据。

Tips

```
单工、半双工和全双工 这三者都是建立在	TCP协议(传输层上)的概念，不要与应用层进行混淆。
```

## 什么是Websocket

Websocket 协议也是基于TCP协议的，是一种双全工通信技术、复用HTTP握手通道。

Websocket默认使用请求协议为:ws:// ,默认端口:80。对TLS加密请求协议为:wss:// ，端口:443。

### 3.1 特点

- 支持浏览器/Nodejs环境

- 支持双向通信

- API简单易用

- 支持二进制传输

- 减少传输数据量

### 3.2 建立连接过程
Websocket复用了HTTP的握手通道。指的是，客户端发送HTTP请求，并在请求头中带上Connection: Upgrade 、Upgrade: websocket，服务端识别该header之后，进行协议升级，使用Websocket协议进行数据通信。

![images](https://github.com/foxliang/Blog/blob/master/images/websocket%E5%8F%82%E6%95%B0.png)

参数说明

- Request URL 请求服务端地址

- Request Method 请求方式 (支持get/post/option)

- Status Code 101 Switching Protocols

RFC 7231 规范定义

规范解释: 当收到101请求状态码时，表明服务端理解并同意客户端请求，更改Upgrade header字段。服务端也必须在response中，生成对应的Upgrade值。



- Connection 设置upgrade header,通知服务端，该request类型需要进行升级为websocket。

- upgrade_mechanism 规范


- Host 服务端 hostname


- Origin 客户端 hostname:port


- Sec-WebSocket-Extensions 客户端向服务端发起请求扩展列表(list)，供服务端选择并在响应中返回


- Sec-WebSocket-Key 秘钥的值是通过规范中定义的算法进行计算得出，因此是不安全的，但是可以阻止一些误操作的websocket请求。


- Sec-WebSocket-Accept

计算公式:  	
```
1. 获取客户端请求header的值: Sec-WebSocket-Key
2. 使用魔数magic = '258EAFA5-E914-47DA-95CA-C5AB0DC85B11'
3. 通过SHA1进行加密计算, sha1(Sec-WebSocket-Key + magic)
4. 将值转换为base64
```

- Sec-WebSocket-Protocol  指定有限使用的Websocket协议，可以是一个协议列表(list)。服务端在response中返回列表中支持的第一个值。


- Sec-WebSocket-Version  指定通信时使用的Websocket协议版本。最新版本:13,历史版本


- Upgrade 通知服务端，指定升级协议类型为websocket

