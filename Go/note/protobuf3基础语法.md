# proto介绍

全称Protocol Buffers(下面简称PB)是Google公司开发的一种数据描述语言，是一种类似 XML 但更灵活和高效的结构化数据存储格式，可用于结构化数据的序列化，适用于数据存储、RPC数据交换格式。它可用于通讯协议、数据存储等领域的语言无关、平台无关、可扩展的序列化结构数据格式。它支持多种语言，比如C++，Java，C#，Python，JavaScript等等。目前它的最新版本是3.3.0。

### proto优点

从上面的proto介绍不难得出其具备下面几个优点：

- 描述简单，对开发人员友好

- 跨平台、跨语言，不依赖于具体运行平台和编程语言

- 高效自动化解析和生成

- 压缩比例高

- 可扩展、兼容性好

### proto3特性

proto3相较于proto2支持更多语言但在语法上更为简洁。去除了一些复杂的语法和特性，更强调约定而弱化语法。

删除原始值字段的presence字段逻辑，删除required字段以及删除默认值。这使得proto3更容易实现如在Android Java，Objective C或 Go 等语言中的开放式结构化表示。

- 移除unknown关键字.

- 去掉extensions类型，使用Any新标准类型替换。

- 针对未知枚举值的固定语法.

- 增加maps(主要指代码生成支持map)

- 添加一组用于表示时间，动态数据等的标准类型。

- 替换二进制编码的明确 JSON 编码

优点：空间效率高，时间效率要高，对于数据大小敏感，传输效率高的

缺点：消息结构可读性不高，序列化后的字节序列为二进制序列不能简单的分析有效性

## 基本结构
```
syntax="proto3";                    //文件第一行指定使用的protobuf版本，如果不指定，默认使用proto2
package services;                   //定义proto包名,可以为.proto文件新增一个可选的package声明符，可选
option go_package = ".;services";     //声明编译成go代码后的package名称，可选的，默认是proto包名

message ProdRequest{                //messaage可以理解为golang中的结构体,可以嵌套
    int32 prod_id=1;                //变量的定义格式为：[修饰符][数据类型][变量名] = [唯一编号] ,同一个message中变量的编号不能相同
}

message ProdResponse{
    int32 pro_stock=1;
}

service ProdService{                                     //定义服务
    rpc GetProdStock (ProdRequest) returns (ProdResponse);  //rpc方法
}
```

## 变量类型

ProtoBuf  | Golang 
---|---
int32/sint32/sfixed32 | int64
int32/sint32/sfixed32 | int32
uint32/fixed32 | uint32
uint64/fixed64 | uint64
float| float32
double| float64
bool| bool
string| string
bytes| []byte
enum| 数组或slice
google.protobuf.Timestamp| timestamp.Timestamp

备注：最后的时间类型golang需要引入包github.com/golang/protobuf/ptypes/timestamp,定义如下
```
t:=timestamp.Timestamp(time.Now().Unix())
```
然后.protp文件需要导入google/protobuf/timestamp.proto

### 修饰符

#### repeated

如果一个字段被repeated修饰，则表示它是一个列表类型的字段，相当于golang里的切片
```
message SearchRequest {
  repeated string args = 1 // 列表类型
}
```
#### reserved
如果你希望可以预留一些数字标签或者字段可以使用reserved修饰符

```
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
  string foo = 3 // 编译报错，因为‘foo’已经被标为保留字段
}
```
### 默认值


类型  | 默认值 
---|---
string| ""
bytes| 空字节
bool| false
数字| 0
enum| 第一个值
message| 跟编程语言有关
repeated| 空列表


### 导入文件

import "myproject/other_protos.proto"; // 这样就可以引用在other_protos.proto文件中定义的message,不能导入不使用的.proto文件

### 定义服务
在你的 .proto 文件中指定 service,然后在service里定义rpc方法即可，要注意指定参数和返回值

```
service RouteGuide {
   rpc GetFeature(Point) returns (Feature) {}
}
```

### rpc方法
gRPC 允许你定义4种类型的 service 方法

### 简单rpc
客户端使用存根发送请求到服务器并等待响应返回，就像平常的函数调用一样

```
service RouteGuide {
   rpc GetFeature(Point) returns (Feature) {}
}
```

### 服务端流式rpc
通过在 响应返回参数 类型前插入 stream 关键字，可以指定一个服务器端的流方法。客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。

```
service RouteGuide {
   rpc ListFeatures(Rectangle) returns (stream Feature) {}
}
```

### 客户端流式rpc
通过在 请求参数 类型前指定 stream 关键字来指定一个客户端的流方法。客户端写入一个消息序列并将其发送到服务器，同样也是使用流。一旦客户端完成写入消息，它等待服务器完成读取返回它的响应。

```
service RouteGuide {
   rpc RecordRoute(stream Point) returns (RouteSummary) {}
}
```

### 双向流式rpc
通过在请求和响应前加 stream 关键字去制定方法的类型。两个流独立操作，因此客户端和服务器可以以任意喜欢的顺序读写：比如， 服务器可以在写入响应前等待接收所有的客户端消息，或者可以交替的读取和写入消息，或者其他读写的组合。
```
service RouteGuide {
   rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
}
```

