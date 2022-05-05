// https://leetcode-cn.com/problems/merge-two-sorted-lists/
package main

import (
	"fmt"
	"leetcode/util"
)

type ListNode = util.ListNode

/**
 * 时间 O(m+n) m: l1 长度 n: l2长度
 * 空间 O(1)
 */
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

/**
 * 优雅地解决头为 nil 判断问题。代码更简洁
 * 时间 O(m+n) m: l1 长度 n: l2长度
 * 空间 O(1)
 */
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	cur := root
	for l1 != nil && l2 != nil {
		var tmp *ListNode
		if l1.Val >= l2.Val {
			tmp = l2
			l2 = l2.Next
		} else {
			tmp = l1
			l1 = l1.Next
		}
		cur.Next = tmp
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return root.Next
}
func main() {
	fmt.Println("vim-go")
}
