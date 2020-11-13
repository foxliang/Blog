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


2.
```
package main

import (
	"fmt"
	"time"
)

// 理念：不要通过共享内存来通信；通过通信来共享内存

/**
channel 例子二： channel数组
*/

func worker2(id int, c chan int) {
	for {
		// 从 channel 里面获取数据，赋值给 n
		n := <-c
		fmt.Println("接收到的数据： ", n)

		// 发现输出的顺序不一样，这是因为，运行的时候是乱序的，谁先执行是随机的

	}
}

func main() {
	// var c chan int // 此时 c == nil
	var channels [10]chan int // channel 数组

	for i := 0; i < 10; i++ {
		// 创建一个 chan
		channels[i] = make(chan int)
		// 接收 channel 的数据，只能在另一个 协程(goruotime)里面才能接收到
		go worker2(i, channels[i])
	}

	// 往  channel 发送数据
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- i
	}

	// 睡眠一下 把两个循环中的值都打印出来
	time.Sleep(time.Millisecond)

}
```
