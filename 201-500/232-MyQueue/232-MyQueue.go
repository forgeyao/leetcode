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

/**
 * 双栈类似实现
 */
type MyQueue2 struct {
	s1, s2 []int
}

/**
 * 时间 O(1)  空间 O(1)
 */
func Constructor() MyQueue2 {
	return MyQueue2{}
}

/**
 * 时间 O(1)
 */
func (this *MyQueue2) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue2) in2out() {
	for len(this.s1) > 0 {
		this.s2 = append(this.s2, this.s1[len(this.s1)-1])
		this.s1 = this.s1[:len(this.s1)-1]
	}
}

/**
 * 时间 O(1)  空间 O(1)
 */
func (this *MyQueue2) Pop() int {
	if len(this.s2) == 0 {
		this.in2out()
	}
	ret := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return ret
}

/**
 * 时间 O(1)  空间 O(1)
 */
func (this *MyQueue2) Peek() int {
	if len(this.s2) == 0 {
		this.in2out()
	}
	return this.s2[len(this.s2)-1]
}

/**
 * 时间 O(1)  空间 O(1)
 */
func (this *MyQueue2) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
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
