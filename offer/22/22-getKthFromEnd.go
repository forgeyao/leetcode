// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
package main

/**
 * 顺序查找。倒数第k个，转换为第 l-k (l为链表长度)
 * 时间 O(N)
 * 空间 O(1)
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	l := getListLength(head)
	if k > l { // 题目没说 k <= l
		k %= l
	}
	ret := head
	for i := 1; i <= l-k; i++ {
		ret = ret.Next
	}
	return ret
}
func getListLength(head *ListNode) int {
	l := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		l++
	}
	return l
}

/**
 * 快慢指针
 * 时间 O(N)
 * 空间 O(1)
 */
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	quick, slow := head, head
	for i := 0; i < k; i++ {
		quick = quick.Next
	}
	for quick != nil {
		quick = quick.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	// test case
	/*
		[] 1
		[1] 1
		[1,2] 1
		[1,2] 2
		[1,2] 3
		[1,2,3,4,5] 2
	*/
}
