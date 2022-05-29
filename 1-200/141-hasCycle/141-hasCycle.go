// https://leetcode.cn/problems/linked-list-cycle/
package main

import (
	"fmt"
	"leetcode/util"
)

type ListNode = util.ListNode

/**
 * 快慢指针
 * 时间 O(N)
 * 空间 O(1)
 */
func hasCycle(head *ListNode) bool {
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		if quick == slow {
			return true
		}
		slow = slow.Next
		quick = quick.Next.Next
	}
	return false
}

func main() {
	forth := new(ListNode)
	forth.Val = -4
	forth.Next = nil

	third := new(ListNode)
	third.Val = 0
	third.Next = forth

	second := new(ListNode)
	second.Val = 2
	second.Next = third

	head := new(ListNode)
	head.Val = 3
	head.Next = second

	//second.Next = head
	//third.Next = second
	forth.Next = second

	fmt.Println("ret:", hasCycle(head))
}
