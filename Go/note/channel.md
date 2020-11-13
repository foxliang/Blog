channel的定义：
```
channel是Go语言中各个并发结构体(goroutine)之前的通信机制。 通俗的讲，就是各个goroutine之间通信的”管道“，有点类似于Linux中的管道。
```
- 1.声明channel

- 2.引用类型
 
- 3.单向channel


var 变量名 chan 数据类型

channel和和map类似，channel也一个对应make创建的底层数据结构的引用。

当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。定义一个channel时，也需要定义发送到channel的值的类型。

// 方法一：channel的创建赋值
```
var ch chan int;

ch = make(chan int);
```
// 方法二：短写法

```
 ch:=make(chan int);
```

// 方法三：综合写法:全局写法！！！！

```
var ch = make(chan int);
```


#### 单向chan

//定义只读的channel

```
read_only := make (<-chan int)
```

 
//定义只写的channel

```
write_only := make (chan<- int)
```
带缓冲区/不带缓冲区 的channel

带缓冲区channel：定义声明时候制定了缓冲区大小(长度)，可以保存多个数据。

ch := make(chan int ,10) //带缓冲区 （只有当队列塞满时发送者会阻塞，队列清空时接受着会阻塞。）

不带缓冲区channel：只能存一个数据，并且只有当该数据被取出时候才能存下一个数据。

ch := make(chan int) //不带缓冲区

#### 无缓冲channel详细解释：

1.一次只能传输一个数据

2.同一时刻，同时有 读、写两端把持 channel，同步通信。

如果只有读端，没有写端，那么 “读端”阻塞。

如果只有写端，没有读端，那么 “写端”阻塞。

读channel： <- channel

写channel： channel <- 数据

#### 举一个形象的例子：

同步通信： 数据发送端，和数据接收端，必须同时在线。 —— 无缓冲channel

打电话。打电话只有等对方接收才会通，要不然只能阻塞

带缓channel详细解释：

#### 举一个形象的例子：

异步通信：数据发送端，发送完数据，立即返回。数据接收端有可能立即读取，也可能延迟处理。 —— 有缓冲channel 不用等对方接受，只需发送过去就行。

发信息。短信。发送完就好，管他什么时候读信息。

#### 如何优雅的关闭channel

注意：

读写操作注意：

```
向已关闭的channel发送数据，则会引发pannic；
channel关闭之后，仍然可以从channel中读取剩余的数据，直到数据全部读取完成。
关闭已经关闭的channel会导致panic
channel如果未关闭，在读取超时会则会引发deadlock异常
```

## select专题：

select是Golang在语言层面提供的多路IO复用的机制，其可以检测多个channel是否ready(即是否可读或可写)

总结select：

```
select语句中除default外，每个case操作一个channel，要么读要么写
select语句中除default外，各case执行顺序是随机的
如果select所有case中的channel都未ready，则执行default中的语句然后退出select流程
select语句中如果没有default语句，则会阻塞等待任一case
select语句中读操作要判断是否成功读取，关闭的channel也可以读取
```
