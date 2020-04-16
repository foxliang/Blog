# goroutine实现并发 

Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。
## sync 进行监工 保证所有worker都执行完
```
package main

import (
	"fmt"
	"sync"
)

//goroutine实现简单并发
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("Hello ", i)
	wg.Done() //通知wg子程序已执行完成
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)   //计数+1
		go hello(i) //开启goroutine执行函数
	}
	fmt.Println("Hello world")
	wg.Wait() //阻塞 等所有程序执行完
}
```
结果：
```
Hello  1
Hello  0
Hello  6
Hello  3
Hello  4
Hello  5
Hello  8
Hello  7
Hello  9
Hello  10
Hello  11
Hello  12
Hello  13
Hello  14
Hello  2
Hello  21
Hello  15
Hello  16
Hello  17
Hello  18
Hello  19
Hello  20
Hello  25
Hello  22
Hello  23
Hello  24
Hello  28
Hello  26
Hello  27
Hello  30
Hello  29
Hello  32
Hello  31
Hello  33
Hello  34
Hello  35
Hello  36
Hello  37
Hello  38
Hello world
Hello  99
Hello  39
Hello  40
Hello  41
Hello  42
Hello  43
Hello  44
Hello  45
Hello  46
Hello  47
Hello  48
Hello  49
Hello  50
Hello  51
Hello  52
Hello  53
Hello  54
Hello  55
Hello  56
Hello  57
Hello  58
Hello  59
Hello  60
Hello  61
Hello  62
Hello  63
Hello  64
Hello  65
Hello  66
Hello  67
Hello  68
Hello  69
Hello  70
Hello  71
Hello  72
Hello  73
Hello  74
Hello  75
Hello  76
Hello  77
Hello  78
Hello  79
Hello  80
Hello  81
Hello  82
Hello  83
Hello  84
Hello  85
Hello  93
Hello  94
Hello  95
Hello  96
Hello  97
Hello  98
Hello  86
Hello  88
Hello  87
Hello  89
Hello  90
Hello  92
Hello  91
```
