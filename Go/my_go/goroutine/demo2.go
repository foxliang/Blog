# GOMAXPROCS
Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	wg.Done()
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(2) //并发时占用cpu的核心
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
```

结果：
```
B: 1
A: 1
A: 2
A: 3
A: 4
A: 5
A: 6
A: 7
A: 8
A: 9
B: 2
B: 3
B: 4
B: 5
B: 6
B: 7
B: 8
B: 9
```
