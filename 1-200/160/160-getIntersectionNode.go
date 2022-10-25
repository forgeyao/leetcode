// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
package main

import (
	"fmt"
	"leetcode/util"
)

type ListNode = util.ListNode

/*
 * 常规思路。让两个链表等长后比较
 * 时间 O(m+n)
 * 空间 O(1)
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 现求出两个链表的长度
	lA, lB := 0, 0
	for index := headA; index != nil; lA, index = lA+1, index.Next {
	}
	for index := headB; index != nil; lB, index = lB+1, index.Next {
	}
	// 较长的链表先移动多出来的长度
	indexA, indexB := headA, headB
	for i := 0; i < lA-lB; i, indexA = i+1, indexA.Next {
	}
	for i := 0; i < lB-lA; i, indexB = i+1, indexB.Next {
	}
	// 从 indexA/indexB 开始，两个链表相同长度，同时遍历
	for ; indexA != indexB; indexA, indexB = indexA.Next, indexB.Next {
	}
	return indexA
}

/**
 * 官方解答,不用计算链表长度，但比较难想到
 * 时间 O(m+n)
 * 空间 O(1)
 */
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	indexA, indexB := headA, headB
	if indexA == nil || indexB == nil {
		return nil
	}
	for indexA != indexB {
		if indexA == nil { // 每轮回一次,长度差值小1,最终变成等长
			indexA = headB
		} else {
			indexA = indexA.Next
		}
		if indexB == nil {
			indexB = headA
		} else {
			indexB = indexB.Next
		}
	}
	return indexA
}

func main() {
	fmt.Println("ret")
}
