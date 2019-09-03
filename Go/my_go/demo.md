## 1.Hello World

```
package main

import "fmt"

func main() { 
    fmt.Println("Hello World")
}
```

输出：
```
Hello World
```

## 2.九九乘法表

```
package main

import (
    "fmt"
)

func main() {
    // 遍历, 决定处理第几行
    for y := 1; y <= 9; y++ {
        // 遍历, 决定这一行有多少列
        for x := 1; x <= y; x++ {
            fmt.Printf("%d*%d=%d ", x, y, x*y)
        }
        // 手动生成回车
        fmt.Println()
    }
}
```

输出：
```
1*1=1 
1*2=2 2*2=4 
1*3=3 2*3=6 3*3=9 
1*4=4 2*4=8 3*4=12 4*4=16 
1*5=5 2*5=10 3*5=15 4*5=20 5*5=25 
1*6=6 2*6=12 3*6=18 4*6=24 5*6=30 6*6=36 
1*7=7 2*7=14 3*7=21 4*7=28 5*7=35 6*7=42 7*7=49 
1*8=8 2*8=16 3*8=24 4*8=32 5*8=40 6*8=48 7*8=56 8*8=64 
1*9=9 2*9=18 3*9=27 4*9=36 5*9=45 6*9=54 7*9=63 8*9=72 9*9=81 
```

## 3.计算函数执行时间

```
package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    test()
    end := time.Now()
    result := end.Sub(start)
    fmt.Printf("该函数执行完成耗时: %s\n", result)
}

func test() {
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum += i
    }
}
```

输出：
```
该函数执行完成耗时: 181.4952ms
```

## 4.搭建一个简单的网站程序

```
package main

import (
    "io"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, world!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
```

运行此程序，并访问 http://127.0.0.1:8080/hello：
```
Hello, world!
```
