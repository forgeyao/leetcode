//https://leetcode-cn.com/problems/reverse-linked-list/
package main

import (
	"fmt"
	"leetcode/util"
)

type ListNode = util.ListNode

/**
 * 常规遍历
 * 时间O(n), n表示链表长度
 * 空间O(1)
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/**
 * 递归版
 * 时间O(n), n表示链表长度
 * 空间O(n)
 */
func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode = nil

	// 为了方便返回头，定义内部递归函数
	var f func(*ListNode) *ListNode
	f = func(h *ListNode) *ListNode {
		if h == nil || h.Next == nil {
			newHead = h
			return h
		}
		prev := f(h.Next)
		prev.Next = h
		h.Next = nil
		return h
	}

	f(head)
	return newHead
}

/*
 * 官方递归版, 更简洁
 * 时间O(n), n表示链表长度
 * 空间O(n)
 */
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	heads := [][]int{{1, 2, 3, 4, 5}, {1, 2}, {}}
	for i := 0; i < len(heads); i++ {
		head := util.CreateListNode(heads[i])
		//newHead := reverseList(head)
		//newHead := reverseList2(head)
		newHead := reverseList3(head)
		for ; newHead != nil; newHead = newHead.Next {
			fmt.Print(newHead.Val, " ")
		}
		fmt.Println()
	}
}
