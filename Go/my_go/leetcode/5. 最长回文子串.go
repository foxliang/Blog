package main

import (
	"fmt"
)

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"


*/

func main() {
	res := longestPalindrome("babad")
	fmt.Println("res", res)
}
func longestPalindrome(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	s = t
	armLen := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var curArmLen int
		if right >= i {
			iSym := j*2 - i
			minArmLen := min(armLen[iSym], right-i)
			curArmLen = expand(s, i-minArmLen, i+minArmLen)
		} else {
			curArmLen = expand(s, i, i)
		}
		armLen = append(armLen, curArmLen)
		if i+curArmLen > right {
			j = i
			right = i + curArmLen
		}
		if curArmLen*2+1 > end-start {
			start = i - curArmLen
			end = i + curArmLen
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//go run main.go
//res aba
