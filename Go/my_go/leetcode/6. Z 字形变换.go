package main

import (
	"fmt"
	"strings"
)

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G


*/

func main() {
	res := convert("LEETCODEISHIRING", 3)
	fmt.Println("res", res)
}
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	n := 2*numRows - 2
	for i, char := range s {
		x := i % n
		rows[min(x, n-x)] += string(char)
	}
	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//go run main.go
//res LCIRETOESIIGEDHN
