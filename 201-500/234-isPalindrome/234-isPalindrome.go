// https://leetcode.cn/problems/palindrome-linked-list/
package main

import (
	"fmt"
	"leetcode/util"
)

type ListNode = util.ListNode

/**
 * 常规方法。用数组缓存遍历结果，再做比较
 * 时间 O(n)
 * 空间 O(n)
 */
func isPalindrome2(head *ListNode) bool {
	v := []int{}
	for node := head; node != nil; node = node.Next {
		v = append(v, node.Val)
	}
	for i, j := 0, len(v)-1; i <= j; i, j = i+1, j-1 {
		if v[i] != v[j] {
			return false
		}
	}
	return true
}

/**
 * 常规方法。用数组缓存遍历结果，再做比较
 * 时间 O(n)
 * 空间 O(1)
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	length := 0
	next := head
	for ; next != nil; next = next.Next {
		length++
	}

	var pre, cur *ListNode
	next = head
	for i := 0; i < length/2; i++ {
		pre = cur
		cur = next
		next = next.Next
		cur.Next = pre
	}

	if length%2 == 1 {
		next = next.Next
	}
	for cur != nil && next != nil {
		if cur.Val != next.Val {
			return false
		}
		cur = cur.Next
		next = next.Next
	}
	return true
}

func main() {
	a := [][]int{{2}, {1, 2}, {1, 2, 1}, {1, 2, 3, 1}, {1, 2, 2, 1}}

	for i := 0; i < len(a); i++ {
		head := new(ListNode)
		cur := head
		for j := 0; j < len(a[i]); j++ {
			node := new(ListNode)
			node.Val = a[i][j]
			node.Next = nil
			if j == 0 {
				head = node
				cur = head
			} else {
				cur.Next = node
				cur = node
			}
			fmt.Print(a[i][j], " ")
		}
		fmt.Println("ret:", isPalindrome(head))
	}
}
