// https://leetcode-cn.com/problems/merge-two-sorted-lists/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var root, cur, tmp *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			tmp = l2
			l2 = l2.Next
		} else {
			tmp = l1
			l1 = l1.Next
		}
		if root == nil {
			root = tmp
			cur = root
		} else {
			cur.Next = tmp
			cur = cur.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return root
}
func main() {
	fmt.Println("vim-go")
}
