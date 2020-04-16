# goroutine实现并发 
sync 进行监工 保证所有worker都执行完
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
