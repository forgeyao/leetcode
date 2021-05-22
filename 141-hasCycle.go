package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	quick := head
	slow := head
	for slow != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if quick == nil {
			return false
		}
		if quick == slow || quick.Next == slow {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, quick := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if quick == slow {
			return true
		}
	}
	return false
}
