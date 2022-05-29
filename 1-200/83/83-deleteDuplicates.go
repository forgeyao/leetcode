// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	for node := head.Next; node != nil; node = node.Next {
		if pre.Val == node.Val {
			pre.Next = node.Next
		} else {
			pre = node
		}
	}
	return head
}

func main() {
	n := [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 1, 2, 3, 3}}
	for i := 0; i < len(n); i++ {
		var root, node *ListNode
		for j := 0; j < len(n[i]); j++ {
			tmp := new(ListNode)
			tmp.Val = n[i][j]
			if j == 0 {
				root = tmp
				node = tmp
			} else {
				node.Next = tmp
				node = tmp
			}
		}
		root = deleteDuplicates(root)
		for ; root != nil; root = root.Next {
			fmt.Print(root.Val, " ")
		}
		fmt.Println()
	}
}
