php7发布已经升级到7.2.里面发生了很多的变化。

参考资料：

http://php.net/manual/zh/migration70.new-features.php

http://php.net/manual/zh/migration71.new-features.php

http://php.net/manual/zh/migration72.new-features.php

# PHP7.0
## PHP7.0新特性
#### 1. 组合比较符 (<=>)
组合比较符号用于比较两个表达式。当$a小于、等于或大于$b时它分别返回-1、0或1，比较规则延续常规比较规则。对象不能进行比较

```
var_dump('PHP' <=> 'Node'); // int(1)
var_dump(123 <=> 456); // int(-1)
var_dump(['a', 'b'] <=> ['a', 'b']); // int(0)
```
#### 2. null合并运算符
由于日常使用中存在大量同时使用三元表达式和isset操作。使用null合并运算符可以简化操作

```
# php7以前
if(isset($_GET['a'])) {
  $a = $_GET['a'];
}
# php7以前
$a = isset($_GET['a']) ? $_GET['a'] : 'none';

# PHP 7
$a = $_GET['a'] ?? 'none';
```
#### 4. 变量类型声明
变量类型声明有两种模式。一种是强制的，和严格的。允许使用下列类型参数int、string、float、bool

同时不能再使用int、string、float、bool作为类的名字了

```
function sumOfInts(int ...$ints)
{
    return array_sum($ints);
}
var_dump(sumOfInts(2, '3', 4.1)); // int(9)

# 严格模式

declare(strict_types=1);

function add(int $x, int $y)
{
    return $x + $y;
}
var_dump(add('2', 3)); // Fatal error: Argument 1 passed to add() must be of the type integer
```
#### 5. 返回值类型声明
增加了返回类型声明，类似参数类型。这样更方便的控制函数的返回值.在函数定义的后面加上:类型名即可

```
function fun(int $a): array
{
  return $a;
}
fun(3);//Fatal error
```

## PHP7 内核架构
![image](https://github.com/weilyf2017/Blog/blob/master/images/PHP7%E5%86%85%E6%A0%B8%E6%9E%B6%E6%9E%84.png)

### zend引擎
词法/语法分析、AST编译和 opcodes 的执行均在 Zend 引擎中实现。此外，PHP的变量设计、内存管理、进程管理等也在引擎层实现。

### PHP层
zend 引擎为 PHP 提供基础能力，而来自外部的交互则需要通过 PHP 层来处理。

### SAPI
server API 的缩写，其中包含了场景的 cli SAPI 和 fpm SAPI。只要遵守定义好的 SAPI 协议，外部模块便可与PHP完成交互。

### 扩展部分
依据 zend 引擎提供的核心能力和接口规范，可以进行开发扩展。

## PHP 7 源码结构
php 7 的源码主要目录有：sapi 、Zend、main、ext 和 TSRM 这几个。

### sapi目录
sapi目录是对输入和输出层的抽象，是PHP提供对外服务的规范。

几种常用的 SAPI：

1）apache2handler: Apache 扩展，编译后生成动态链接库，配置到Apache下。当有 http 请求到 Apache 时，根据配置会调用此动态链接库来执行PHP代码，完成与PHP的交互。

2）cgi-fcgi: 编译后生成支持 CGI 协议的可执行程序，webserver（如NGINX）通过 CGI 协议把请求传给CGI进程，CGI 进程根据请求执行相应代码后将执行结果返回给 webserver。

3）fpm-fcgi: fpm是 FastCGI 进程管理器。以 NGINX 服务器为例，当有请求发送到 NGINX 服务器，NGINX 按照 FastCGI 协议把请求交给 php-fpm 进程处理。

4）cli: PHP的命令行交互接口

### Zend 目录
Zend 目录是 PHP 的核心代码。PHP中的内存管理，垃圾回收、进程管理、变量、数组实现等均在该目录的源码里。

### main 目录
main目录是SAPI层和Zend层的黏合剂。Zend 层实现了 PHP 脚本的编译和执行，sapi 层实现了输入和输出的抽象，main目录则在它们中间起着承上启下的作用。承上，解析 SAPI 的请求，分析要执行的脚本文件和参数；启下，调用 zend 引擎之前，完成必要的模块初始化等工作。

### ext目录
ext 是 PHP 扩展相关的目录，常用的 array、str、pdo 等系列函数都在这里定义。

### TSRM
TSRM（Thread Safe Resource Manager）——线程安全资源管理器， 是用来保证资源共享的安全。

参考 
《PHP7 底层设计与源码解析》
