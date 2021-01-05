## 1.protobuf3 简介 安装

[基础语法](https://github.com/foxliang/Blog/blob/master/Go/note/protobuf3%E5%9F%BA%E7%A1%80%E8%AF%AD%E6%B3%95.md)

使用protobuf-go
```
go get github.com/golang/protobuf/proto
go get github.com/gogo/protobuf/gogoproto
go get github.com/gogo/protobuf/protoc-gen-gogo
go get github.com/gogo/protobuf/protoc-gen-gofast
```

## 2.写test.proto文件

相当于给一个对象添加相应的属性。

```
//指定版本
//注意proto3与proto2的写法有些不同
syntax = "proto3";

//包名，通过protoc生成时go文件时
package student;

// 班级
message Class {
    int32 num = 1;
    repeated Student students = 2;
}

// 学生
message Student {
    string name = 1;
    int32 age = 2;
    Sex sex = 3;
}

//性别
enum Sex {
    MAN = 0;
    WOMAN = 1;
}
```

## 3.生成文件test.pb.go文件

.proto文件写好之后，不方便我们在代码中使用，需要利用刚才安装的proto工具生成一个我们可以在代码中方便实际调用的类。

这个类生成之后就变成我们和protobuf交换数据的桥梁，我们可以看懂和使用，protobuf也可以识别和解析。

生成test.pb.go文件之后.proto就不需要了，但是为了后期更改和代码可读性继续保留该文件。

test.pb.go具体代码我就不贴出来了，命令如下： 
```
protoc --go_out=. *.proto
```
## 4.测试和验证

go run main.go

```
2021/01/05 12:00:10 1
2021/01/05 12:00:10 小明 21 MAN
2021/01/05 12:00:10 小花 22 WOMAN
2021/01/05 12:00:10 小牛 20 MAN
```
