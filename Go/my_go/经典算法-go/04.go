package main

import (
	"fmt"
)

func getPositiveInt() uint64 {
	var n uint64

	fmt.Println("请输入一个正整数：")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	return n
}

func main() {
	num := int(getPositiveInt())

	fmt.Printf("%d=", num)
	for {
		var i int
		for i = 2; i <= num; i++ {
			if num%i == 0 {
				if num != i {
					fmt.Printf("%d*", i)
				}
				num = num / i
				break
			}
		}

		if num == 1 {
			fmt.Println(i)
			break
		}
	}
}
