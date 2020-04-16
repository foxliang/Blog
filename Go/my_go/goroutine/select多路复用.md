# select的使用类似于switch语句，它有一系列case分支和一个默认的分支。

每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句

举个小例子来演示下select的使用：
```
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
```

结果：
```
0
2
4
6
8
```
使用select语句能提高代码的可读性。

- 可处理一个或多个channel的发送/接收操作。
- 如果多个case同时满足，select会随机选择一个。
- 对于没有case的select{}会一直等待，可用于阻塞main函数。
