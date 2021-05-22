package main

import "fmt"

func main() {
	/*var i int
	fmt.Scanf("%d", &i)
	fmt.Println("input:", i)
	*/
	//fmt.Println("result:", reverse2(i))
	//fmt.Println("result:", isPalindrome(i))

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
