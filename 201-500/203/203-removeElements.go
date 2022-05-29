// https://leetcode-cn.com/problems/remove-linked-list-elements/
package main

import (
	"leetcode/util"
)

type ListNode = util.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	var pre *ListNode
	for h := head; h != nil; h = h.Next {
		if h.Val == val {
			if pre == nil {
				head = head.Next
			} else {
				pre.Next = h.Next
			}
		} else {
			pre = h
		}
	}
	return head
}

//  优化版. 解决头节点删除
func removeElements2(head *ListNode, val int) *ListNode {
	var newHead *ListNode = &ListNode{Next: head}
	for h := newHead; h.Next != nil; {
		if h.Next.Val == val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	return newHead.Next
}

func main() {
	n := [][]int{{1, 2, 6, 3, 4, 5, 6}, {}, {7, 7, 7, 7}}
	k := []int{6, 1, 7}
	for i := 0; i < len(n) && i < len(k); i++ {
		h := removeElements(util.CreateListNode(n[i]), k[i])
		h2 := removeElements2(util.CreateListNode(n[i]), k[i])
		util.PrintList(h)
		util.PrintList(h2)
	}
}
