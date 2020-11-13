1.

```
package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for {
		// 从 channel 里面获取数据，赋值给 n
		n := <-c
		fmt.Println("接收到的数据： ", n)
	}
}

func main() {
	// var c chan int // 此时 c == nil

	// 创建一个 chan
	c := make(chan int)

	// 接收 channel 的数据，只能在另一个 协程(goruotime)里面才能接收到
	go worker(c)

	// 往  channel 发送数据
	c <- 1
	c <- 2

	// 睡眠一下
	// 因为只能打印出 1，还来不及打印 2 的时候，程序已经结束了
	// 所以睡眠一下 保证1和2 都能打印出来
	time.Sleep(time.Millisecond)

}

```
