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
### run
```
easyjson -all <file>.go
```
