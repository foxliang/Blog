package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 101; i <= 200; i++ {
		var prime = true
		ul := math.Sqrt(float64(i + 1))
		for j := 2; j < int(ul); j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}

		if prime {
			fmt.Println(i)
		}
	}
}
