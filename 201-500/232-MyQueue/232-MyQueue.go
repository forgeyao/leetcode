// https://leetcode-cn.com/problems/implement-queue-using-stacks/
package main

import (
	"fmt"
	"leetcode/util"
)

type MyQueue struct {
	s1 util.Stack // 用 slice 封装了 Stack
	s2 util.Stack
}

/**
 * 时间 O(1)  空间 O(1)
 */
func Constructor() MyQueue {
	return MyQueue{s1: util.Constructor(), s2: util.Constructor()}
}

/**
 * 时间 O(n)  空间 O(n)
 */
func (this *MyQueue) Push(x int) {
	for i := this.s1.Size() - 1; i >= 0; i-- {
		this.s2.Push(this.s1.Pop())
	}
	this.s1.Push(x)
	for i := this.s2.Size() - 1; i >= 0; i-- {
		this.s1.Push(this.s2.Pop())
	}
}

/**
 * 时间 O(1)  空间 O(1)
 */
func (this *MyQueue) Pop() int {
	return this.s1.Pop()
}

/**
 * 时间 O(1)  空间 O(1)
 */
func (this *MyQueue) Peek() int {
	return this.s1.Peek()
}

/**
 * 时间 O(1)  空间 O(1)
 */
func (this *MyQueue) Empty() bool {
	return this.s1.Empty()
}

func main() {
	cmds := []string{"MyQueue", "push", "push", "peek", "pop", "empty"}
	nums := []int{0, 1, 2, 0, 0, 0}
	var q MyQueue
	for i := 0; i < len(cmds) && i < len(nums); i++ {
		switch cmds[i] {
		case "MyQueue":
			q = Constructor()
			fmt.Println("Constructor")
		case "push":
			q.Push(nums[i])
			fmt.Println("push:", q.s1)
		case "peek":
			fmt.Println("peek:", q.Peek())
		case "pop":
			fmt.Println("pop:", q.Pop(), "queue:", q.s1)
		case "empty":
			fmt.Println("empty:", q.Empty())
		}
	}
}
