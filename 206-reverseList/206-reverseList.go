//https://leetcode-cn.com/problems/reverse-linked-list/
package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre, in, post *ListNode
	pre, in, post = nil, head, head.Next
	for in != nil {
		in.Next = pre
		pre = in
		in = post
		if post == nil {
			break
		}
		post = post.Next
	}
	return pre
}
func main() {
	heads := [][]int{{1, 2, 3, 4, 5}, {1, 2}, {}}
	for i := 0; i < len(heads); i++ {
		var head, index *ListNode
		for j := 0; j < len(heads[i]); j++ {
			if j == 0 {
				head = new(ListNode)
				head.Val = heads[i][j]
				head.Next = nil
				index = head
			} else {
				data := new(ListNode)
				data.Val = heads[i][j]
				data.Next = nil
				index.Next = data
				index = index.Next
			}
		}

		newHead := reverseList(head)
		for ; newHead != nil; newHead = newHead.Next {
			fmt.Print(newHead.Val, " ")
		}
		fmt.Println()
	}
}
