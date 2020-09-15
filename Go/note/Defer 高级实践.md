defer 是一个用起来非常简单的特性。

它的实现原理也不复杂。


本文主要介绍这个特性在实际项目中的利弊以及建议。

## 为什么要用 defer

任何一个特性都有它的设计初衷，主要是被用来解决什么问题的，任何一个特性也都有它合适和不合适出现的地方，我们清楚地了解并正确合理地使用，是非常重要的。

### 优势

提高安全性、健壮性

让代码更优雅

### 劣势

可读性、可维护性

（注意：用 defer 当然肯定比不用有一定的性能开销，但我们可以忽略，因为影响确实很小。 换句话说，绝大部分情况下，考虑是否使用 defer 时，性能开销不应该是首先考虑的因素。但是！如果你的代码是微秒级别的，那还是要评估后再使用）

### defer 怎么用

官方文档，告诉你 defer 的基本用法

几乎所有其他文章里说 defer 如何如何有坑，defer 需要注意什么等等。。都是官方文档上讲到的三点，在此就不赘述了。下面我分成三部分，建议使用、中立和不建议。

建议使用 是官方 src 里都在用的，而且也是 defer 的设计初衷。

中立 是工程实践中总结出来，平衡了代码优雅和可读性、可维护性后的结果。

不建议 是弊大于利，得不偿失的用法，主要影响的就是降低可读性，可维护性。

### 建议使用

#### Recover
```
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
    }
}()
```
资源回收

各种资源的使用，如果在用完之后不 close，就会造成资源的泄露，可能会严重影响程序运行，甚至造成程序死掉

#### 网络 I/O
```
c, err := Dial("udp", raddr)
if err != nil {
    return err
}
defer c.Close()
文件 I/O
f, err := os.Open(filename)
if err != nil {
    return
}
defer f.Close()
channel 关闭
fd, _ := os.Open("txt")
errc := make(chan error, 1)
// 主动关闭，减小 GC 压力。
defer close(errc)
    
var buf [1]byte
n, err := fd.Read(buf[:1])
if n == 0 || err != nil {
    errc <- fmt.Errorf("read byte = %d, err = %v", n, err)
}
```
#### 避免死锁
```
type A struct {
    t int
    sync.Mutex
}

func main() {
    a := new(A)
    for i := 0; i < 2000; i++ {
        go a.incr()
    }
    time.Sleep(500 * time.Millisecond) // 此处用 sleep 简单模拟等待同步，实际这样写不严谨，可用 waitGroup、channel 等
    fmt.Println(a.t)
}

func (a *A) incr() {
    a.Lock()
    defer a.Unlock()
    
    // 模拟 ... 一堆逻辑

    // 然后 ... 中间有好几个 return 出口
    
    // 如果我们不用 defer，就要在每个 return 都写上 a.Unlock，不然就可能会造成死锁    
    a.t++
}
```
### 中立

函数返回时的打点

记日志

这里可能稍微有一些复杂，我稍微讲一下

第一步，会先执行 log("do") 调用 log 函数传入参数 “do”

第二步，log 函数执行函数体即 start := time.Now() fmt.Printf("enter %s\n", msg)两行，然后给调用方 do 函数返回一个 func()

第三步，这个 func() 被放到 defer 里，等到 do 函数返回时才会执行。
```
func main() {
    do()
}

func do() {
    defer log("do")()

    // ... 一些逻辑

    time.Sleep(1 * time.Second)
}

func log(msg string) func() {
    start := time.Now()
    fmt.Printf("enter %s\n", msg)
    return func() { fmt.Printf("exit %s (%s)", msg, time.Since(start)) }
}
```
#### 错误处理
因为 go 自带的比较恶心的 err != nil 的判断，业务逻辑中可能会有大量的这种代码，而我们又要对出错进行一个统一的处理的时候，可以用。

数据库事务的回滚操作
```
tx, err := db.Begin()
if err != nil {
    return err
}
defer func() {
    if err != nil {
        tx.Rollback()
    }
}()
```
// ... 中间会发生多个数据库操作 ...

// 提交，那么在提交之前发生的任何错误，返回时都可利用之前注册的 defer 进行回滚
tx.Commit()

#### 不建议
不建议的用法就不给出代码示例了，怕你看了错误的代码示例反而记住了，就不好了。下面只说不建议的用法场景。

##### 不要直接在循环中使用 defer

defer 是后定义的先执行，和栈类似。

如果在循环中调用 defer，可能会导致堆积了很多 defer，在循环结束后才会执行。

这中间如果有任何一个 defer 失败了怎么办？

多个 defer 执行的内容有没有依赖关系和冲突？

所以，除非万不得已，不要给自己增加复杂度。

不这么用就好了。

##### 不要在 defer 中传入体积很大的参数

因为编译器的很多优化对它都不起作用，所以尽量不要传入体积很大的参数，当然我觉得也应该没有多少人会传入一堆参数来用 defer 的。

##### 不要用 receiver 调用 defer
因为 receiver 是当做第一个参数传给调用函数的，也是值传递，除非你能时刻明确注意 receiver 是否是一个指针，否则最好不要用 defer，不然可能无法得到你想要的结果。

##### 未完待续。。。
defer 原理简述

defer 源码实现的位置：runtime/panic.go

看到这知道我在建议使用中第一个就写 recover 是为什么了吧。

这个特性最初的目的就是给 recover 用的。

编译器会把 defer 关键字转化为对此函数的调用：

func deferproc(siz int32, fn *funcval)
然后当原函数 return 时，会调用：

func deferreturn(arg0 uintptr)
看，它只有一个参数，就是 arg0，也就是 代码中 defer 后面跟着的函数。明显的，只有函数体本身会延迟执行，函数的参数在注册 defer 之前就已经执行完了。

结语

老老实实写代码，不要总想玩魔法。
