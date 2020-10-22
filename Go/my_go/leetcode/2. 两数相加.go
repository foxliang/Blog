package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	tmp1 := l1
	tmp2 := l2
	nums1 := []int{2, 4, 3}
	nums2 := []int{5, 6, 4}

	for i := 0; i < len(nums1); i++ {
		tmp1.Val = nums1[i]
		if i == len(nums1)-1 {
			break
		}
		tmp1.Next = new(ListNode)
		tmp1 = tmp1.Next
	}
	for i := 0; i < len(nums2); i++ {
		tmp2.Val = nums2[i]
		if i == len(nums2)-1 {
			break
		}
		tmp2.Next = new(ListNode)
		tmp2 = tmp2.Next
	}

	fmt.Println("l2", l2.getInt())
	fmt.Println("l1", l1.getInt())
	l2 = addTwoNumbers(l1, l2)
	fmt.Println("res", l2.getInt())
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

// 返回链表中的值
func (l *ListNode) getInt() int {
	if l == nil {
		return 0
	}
	var nums []int
	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	rst := 0
	for k, v := range nums {
		rst += pow10(k) * v
	}
	return rst
}

// 返回10的几次方 <---- 大数字的情况下在这里首先溢出, 写在这里只为了测试用
func pow10(a int) int {
	if a == 0 {
		return 1
	}
	rst := 1
	for i := 1; i <= a; i++ {
		rst *= 10
	}
	return rst
}


//go run main.go
//l2 465
//l1 342
//res 807
