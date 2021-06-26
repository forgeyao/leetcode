package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(n []int) *ListNode {
	var head, index *ListNode
	for i := 0; i < len(n); i++ {
		if i == 0 {
			head = new(ListNode)
			head.Val = n[i]
			head.Next = nil
			index = head
		} else {
			data := new(ListNode)
			data.Val = n[i]
			data.Next = nil
			index.Next = data
			index = index.Next
		}
	}
	return head
}
