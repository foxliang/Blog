package main

import (
	"fmt"
        "time"
)

func main() {
	ti := time.Now()
	fmt.Println("开始 : ", ti.Format("2006-01-02 15:04:05"))
	for i := 1; i<=100; i++{
		fmt.Printf("%d%% [%s]\r",i,getS(i,"#") + getS(100-i," "))
		time.Sleep(time.Duration(100) * time.Millisecond)  //延迟输出
	}
	elapsed := time.Since(ti)
	fmt.Println("\n结束，总共耗时: ", elapsed)
}


func getS(n int,char string) (s string) {
    for i:=1;i<=n;i++{
        s+=char
    }
    return
}


