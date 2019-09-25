## debug_backtrace

[debug_backtrace](https://www.php.net/manual/en/function.debug-backtrace.php) 函数生成 backtrace。

该函数显示由 debug_backtrace() 函数代码生成的数据。

返回一个关联数组。下面是可能返回的元素：

名称 | 类型 | 描述
---|---|---
function | string | 当前的函数名。
line | integer | 当前的行号。
file | string | 当前的文件名。
class | string | 当前的类名。
object | object | 当前对象。
type | string | 当前的调用类型，可能的调用：返回："->" - 方法调用,返回："::" - 静态方法调用,返回 nothing - 函数调用
args | array | 如果在函数中，列出函数参数。如果在被引用的文件中，列出被引用的文件名。
