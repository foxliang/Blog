本文主要讲实践，原理部分会一笔带过，关于 go 语言并发实现和内存模型后续会有文章。
channel 实现的源码不复杂，推荐阅读，https://github.com/golang/go/...

## channel 是干什么的
意义：channel 是用来通信的
实际上：（数据拷贝了一份，并通过 channel 传递，本质就是个队列）

## channel 应该用在什么地方
核心：需要通信的地方
例如以下场景：

-通知广播

-交换数据

-显式同步

-并发控制

...
记住！channel 不是用来实现锁机制的，虽然有些地方可以用它来实现类似读写锁，保护临界区的功能，但不要这么用！

## channel 用例实现

### 超时控制
```
// 利用 time.After 实现
func main() {
    done := do()
    select {
    case <-done:
        // logic
    case <-time.After(3 * time.Second):
        // timeout
    }
}

func do() <-chan struct{} {
    done := make(chan struct{}, 1)
    go func() {
        // do something
        // ...
        done <- struct{}{}
    }()
    return done
}
```
### 取最快的结果

比较常见的一个场景是重试，第一个请求在指定超时时间内没有返回结果，这时重试第二次，取两次中最快返回的结果使用。

超时控制在上面有，下面代码部分就简单实现调用多次了。
```
func main() {
    ret := make(chan string, 3)
    for i := 0; i < cap(ret); i++ {
        go call(ret)
    }
        fmt.Println(<-ret)
}

func call(ret chan<- string) {
    // do something
    // ...
    ret <- "result"
}
```
### 限制最大并发数
// 最大并发数为 2
```
limits := make(chan struct{}, 2)
for i := 0; i < 10; i++ {
    go func() {
        // 缓冲区满了就会阻塞在这
        limits <- struct{}{}
        do()
        <-limits
    }()
}
```
### for...range 优先
for ... range c { do } 这种写法相当于 if _, ok := <-c; ok { do }
```
func main() {
    c := make(chan int, 20)
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c)
    }()
    // 当 c 被关闭后，取完里面的元素就会跳出循环
    for x := range c {
        fmt.Println(x)
    }
}
```
### 多个 goroutine 同步响应
利用 close 广播
```
func main() {
    c := make(chan struct{})
    for i := 0; i < 5; i++ {
        go do(c)
    }
    close(c)
}

func do(c <-chan struct{}) {
    // 会阻塞直到收到 close
    <-c
    fmt.Println("hello")
}
```
### 非阻塞的 select
select 本身是阻塞的，当所有分支都不满足就会一直阻塞，如果想不阻塞，那么一个什么都不干的 default 分支是最好的选择
```
select {
case <-done:
    return
default:   
}
for{select{}} 终止
尽量不要用 break label 形式，而是把终止循环的条件放到 for 条件里来实现

for ok {
    select {
    case ch <- 0:
    case <-done:
        ok = false
    }
}
```
未完待续
...

### channel 特性

基础特性

操作	值为 nil 的 channel	被关闭的 channel	正常的 channel

close	panic	panic	成功关闭

c<-	永远阻塞	panic	阻塞或成功发送

<-c	永远阻塞	永远不阻塞	阻塞或成功接收

### happens-before 特性

无缓冲时，接收 happens-before 发送

任何情况下，发送 happens-before 接收

close happens-before 接收

参考:

https://go101.org/article/channel.html

https://golang.org/doc/effective_go.html#channels
