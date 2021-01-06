# easyjson是什么呢？

根据官网介绍，easyjson是提供高效快速且易用的结构体structs<-->json转换包。easyjson并没有使用反射方式实现，所以性能比其他的json包该4-5倍，比golang 自带的json包快2-3倍。 easyjson目标是维持生成去代码简单，以致于它可以轻松地进行优化或固定。

### 安装
```
go get -u github.com/mailru/easyjson/
go install  github.com/mailru/easyjson/easyjson
or
go build -o easyjson github.com/mailru/easyjson/easyjson
```
### 验证是否安装成功。
```
easyjson
Usage of easyjson:
  -all
        generate marshaler/unmarshalers for all structs in a file
  -build_tags string
        build tags to add to generated file
  -byte
        use simple bytes instead of Base64Bytes for slice of bytes
  -disable_members_unescape
        don't perform unescaping of member names to improve performance
  -disallow_unknown_fields
        return error if any unknown field in json appeared
  -gen_build_flags string
        build flags when running the generator while bootstrapping
  -leave_temps
        do not delete temporary files
  -lower_camel_case
        use lowerCamelCase names instead of CamelCase by default
  -no_std_marshalers
        don't generate MarshalJSON/UnmarshalJSON funcs
  -noformat
        do not run 'gofmt -w' on output file
  -omit_empty
        omit empty fields by default
  -output_filename string
        specify the filename of the output
  -pkg
        process the whole package instead of just the given file
  -snake_case
        use snake_case names instead of CamelCase by default
  -stubs
        only generate stubs for marshaler/unmarshaler funcs

```
其中有几个选项需要注意：

```
-lower_camel_case:将结构体字段field首字母改为小写。如Name=>name。  
-build_tags string:将指定的string生成到生成的go文件头部。  
-no_std_marshalers：不为结构体生成MarshalJSON/UnmarshalJSON函数。  
-omit_empty:没有赋值的field可以不生成到json，否则field为该字段类型的默认值。
-output_filename:定义生成的文件名称。
-pkg:对包内指定有`//easyjson:json`结构体生成对应的easyjson配置。
-snke_case:可以下划线的field如`Name_Student`改为`name_student`。
```
### 使用
记得在需要使用easyjson的结构体上加上//easyjson:json。 如下：

```
package proto

import "time"

//easyjson:easyjson
type School struct {
	Name string `easyjson:"name"`
	Addr string `easyjson:"addr"`
}

//easyjson:easyjson
type Student struct {
	Id       int       `easyjson:"id"`
	Name     string    `easyjson:"s_name"`
	School   School    `easyjson:"s_chool"`
	Birthday time.Time `easyjson:"birthday"`
}
```
### 在结构体包下执行
```
easyjson  -all student.go
```
此时在该目录下出现一个新的文件。

### 现在可以写一个测试类啦。

```
package main

import (
	easyjson "easyjson/proto"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//
	s := easyjson.Student{
		Id:   11,
		Name: "qq",
		School: easyjson.School{
			Name: "CUMT",
			Addr: "xz",
		},
		Birthday: time.Now(),
	}
	bt, err := s.MarshalJSON()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("bt", string(bt))
	//json := `{"id":11,"name":"s_name","s_chool":{"name":"CUMT","addr":"xz"},"birthday":"2017-08-04T20:58:07.9894603+08:00"}`
	ss := easyjson.Student{}
	ss.UnmarshalJSON(bt)
	fmt.Println("ss", ss)

	end := time.Since(start)
	fmt.Println("Since", end)
}
```

运行结果：

```
bt {"Id":11,"Name":"qq","School":{"Name":"CUMT","Addr":"xz"},"Birthday":"2021-01-06T13:39:34.162613917+08:00"}
ss {11 qq {CUMT xz} 2021-01-06 13:39:34.162613917 +0800 CST}
Since 116.135µs
```
