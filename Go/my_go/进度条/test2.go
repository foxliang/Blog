package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(os.Stdout, "result is %d\r", i)
		time.Sleep(time.Duration(600) * time.Millisecond)
	}
	fmt.Println("over")
}

//overlt is 9
