- runtime.Gosched()

用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其它等待的任务运行，并在下次某个时候从该位置恢复执行。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。


- runtime.Goexit()

调用此函数会立即使当前的goroutine的运行终止（终止协程），而其它的goroutine并不会受此影响。runtime.Goexit在终止当前goroutine前会先执行此goroutine的还未执行的defer语句。请注意千万别在主函数调用runtime.Goexit，因为会引发panic。


- runtime.GOMAXPROCS()

用来设置可以并行计算的CPU核数最大值，并返回之前的值。

默认此函数的值与ＣＰＵ逻辑个数相同，即有多少个goroutine并发执行，当然可以设置它，它的取值是１～２５６。最好在主函数在开始前设置它，因为设置它会停止当前程序的运行。

GO默认是使用一个CPU核的，除非设置runtime.GOMAXPROCS

那么在多核环境下，什么情况下设置runtime.GOMAXPROCS会比较好的提高速度呢？

适合于CPU密集型、并行度比较高的情景。如果是IO密集型，CPU之间的切换也会带来性能的损失。


#### Gosched()代码案例

①：没有使用Gosched函数

```
package main
import (
    "fmt"
)

func main() {

    go func() { //子协程   //没来的及执行主进程结束
        for i := 0; i < 5; i++ {
            fmt.Println("go")
        }
    }()

    for i := 0; i < 2; i++ { //默认先执行主进程主进程执行完毕
        fmt.Println("hello")
    }
}
```

结果：

```
hello
hello
```

②：使用Gosched函数

```
package main
import (
    "fmt"
    "runtime"
)

func main() {
    go func() {  //让子协程先执行
        for i := 0; i < 5; i++ {
            fmt.Println("go")
        }
    }()

    for i := 0; i < 2; i++ {
        //让出时间片，先让别的协议执行，它执行完，再回来执行此协程
        runtime.Gosched()
        fmt.Println("hello")
    }
}
```

结果：

```
go
go
go
go
go
hello
hello
```
#### Goexit()代码案例


```
package main

import (
    "fmt"
    "runtime"
)

func test() {
    defer fmt.Println("ccccccccccccc")

    //return //终止此函数
    runtime.Goexit() //终止所在的协程

    fmt.Println("dddddddddddddddddddddd")
}

func main() {

    //创建新建的协程
    go func() {
        fmt.Println("aaaaaaaaaaaaaaaaaa")

        //调用了别的函数
        test()

        fmt.Println("bbbbbbbbbbbbbbbbbbb")
    }() //别忘了()

    //特地写一个死循环，目的不让主协程结束
    for {
    }
}
```
结果：
```
aaaaaaaaaaaaaaaaaa
ccccccccccccc
```

#### GOMAXPROCS()代码案例


```
package main

import (
    "fmt"
    "runtime"
)

func main() {
    //n := runtime.GOMAXPROCS(1) //指定以1核运算
    n := runtime.GOMAXPROCS(4) //指定以4核运算
    fmt.Println("n = ", n)

    for {
        go fmt.Print(1)
        fmt.Print(0)
    }
}
```
