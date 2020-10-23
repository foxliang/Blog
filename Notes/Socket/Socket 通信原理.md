## 一、Socket通信简介

与服务器的通信方式主要有两种：

- Http通信

- Socket通信

两者的最大差异在于：

Http连接使用的是“请求-响应方式”，即在请求时建立连接通道，当客户端向服务器发送请求后，服务端才能向客户端返回数据。

Socket通信则是在双方建立连接后，可以直接进行数据的传输，在连接时可实现信息的主动推送，而不需要每次由客户端向服务器发送请求。

那么，什么是socket？

socket又称套接字，在程序内部提供了与外界通信的端口，即端口通信。

通过建立socket连接，可为通信双方的数据传输提供通道。socket的主要特点有数据丢失率低，使用简单且易于移植。

### 1、什么是Socket

socket是一种抽象层，应用程序通过它来发送和接受数据，使用Socket可以将应用程序添加到网络中，与处于同一网络中的其他应用程序进行通信。

简单来说，Socket提供了程序内部与外界通信的端口并为通信双方提供数据传输通道。

### 2、Socket分类根据不同的底层协议，Socket的实现是多样化的。

在这主要介绍TCP/IP协议簇当中主要的Socket类型为流套接字（streamsocket）和数据报套接字（datagramsocket）。

流套接字将TCP作为其端对端协议，提供了一个可信赖的字节流服务。数据报嵌套字使用UDP协议，提供数据打包发送数据。

## 二、Socket通信流程

![images](https://github.com/foxliang/Blog/blob/master/images/socket%E9%80%9A%E4%BF%A1%E5%8E%9F%E7%90%86.png)



## 三、Socket基本实现原理

#### 1、基于TCP协议的Socket服务端

首先声明一个ServerSocket对象并且指定端口号，然后调用Serversocket的accept()方法接受客户端的数据

。accept()方法在没有数据进行接受时处于堵塞状态。

（Socket socket = serversocket.accept()），一旦接受数据，通过inputstream读取接受的数据。

客户端创建一个Socket对象，执行服务器端的ip地址和端口号（Socket socket = new Socket("172.168.10.108", 8080);），通过inputstream读取数据，获取服务器发出的数据（OutputStream outputstream = socket.getOutputStream();），最后将要发送的数据写入到outputstream即可进行TCP协议的socket数据传输。

#### 2、基于UDP协议的数据传输服务器端首先创建一个DatagramSocket对象，并且指定监听端口。

接下来创建一个空的DatagramSocket对象用于接收数据（byte data[] = new byte[1024]; DatagramSocket packet = new DatagramSocket(data, data.length);），使用DatagramSocket的receive()方法接受客户端发送的数据，receive()与serversocket的accept()方法类似，在没有数据进行接受时处于堵塞状态。

客户端也创建个DatagramSocket对象，并且指定监听的端口。接下来创建一个InetAddress对象，这个对象类似于一个网络的发送地址（InetAddress serveraddress = InetAddress.getByName("172.168.1.120")）。

定义要发送的一个字符串，创建一个DatagramPacket对象，并指定要将该数据包发送到网络对应的那个地址和端口号，最后使用DatagramSocket的对象的send()发送数据。
```
（String str = "hello"; byte data[] = str.getByte(); DatagramPacket packet = new DatagramPacket(data, data.length, serveraddress, 4567）; socket.send(packet);）
```
总结：

- TCP使用的是流的方式发送

- UDP是以包的形式发送

## 四、连接保活

我想你不难发现一个问题，那就是当socket连接成功建立后，如果中途发生异常导致其中一方断开连接，此时另一方是无法发现的，只有在再次尝试发送/接收消息才会因为抛出异常而退出。

简单的说，就是我们维持的socket连接，是一个长连接，但我们没有保证它的时效性，上一秒它可能还是可以用的，但是下一秒就不一定了。

#### 4.1 使用心跳包

保证连接随时可用的最常见方法就是定时发送心跳包，来检测连接是否正常。这对于实时性要求很高的服务而言，还是非常重要的（比如消息推送）。

大体的方案如下：
```
双方约定好心跳包的格式，要能够区别于普通的消息。
客户端每隔一定时间，就向服务端发送一个心跳包
服务端每接收到心跳包时，将其抛弃
如果客户端的某个心跳包发送失败，就可以判断连接已经断开
如果对实时性要求很高，服务端也可以定时检查客户端发送心跳包的频率，如果超过一定时间没有发送可以认为连接已经断开
```
#### 4.2 断开时重连

使用心跳包必然会增加带宽和性能的负担，对于普通的应用我们其实并没有必要使用这种方案，如果消息发送时抛出了连接异常，直接尝试重新连接就好了。

跟上面的方案对比，其实这个抛出异常的消息就充当了心跳包的角色。

总的来说，连接是否要保活，如何保活，需要根据具体的业务场景灵活地思考和定制。

参考：

https://juejin.im/post/6844903593808494600#heading-8

https://juejin.im/post/6844904022567026695#heading-6
