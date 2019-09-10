package main

import (
	"fmt"
	"math"
)

func resolve(n int) (int, int, int) {
	var hundreds, tens, units int

	hundreds = n / 100
	tens = n % 100 / 10
	units = n % 10

	return hundreds, tens, units
}

func judge(n int) bool {
	h, t, u := resolve(n)
	if int(math.Pow(float64(h), 3.0)+math.Pow(float64(t), 3.0)+math.Pow(float64(u), 3.0)) == n {
		return true
	}

	return false
}

func main() {
	for i := 100; i <= 999; i++ {
		if judge(i) {
			fmt.Println(i)
		}
	}
}
