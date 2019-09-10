package main

import (
	"fmt"
)

func main() {
	var rabbit []int64

	rabbit = append(rabbit, 1)
	rabbit = append(rabbit, 1)

	for i := 1; i <= 40; i++ {
		if i > 2 {
			rabbit = append(rabbit, rabbit[i-2]+rabbit[i-3])
		}

		fmt.Printf("第%2d月，有%10d对兔子。\n", i, rabbit[i-1])
	}
}
