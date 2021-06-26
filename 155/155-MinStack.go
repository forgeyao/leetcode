// https://leetcode-cn.com/problems/min-stack/
package main

import "fmt"

type MinStack struct {
	Stack  []int
	MinIdx int // 题目只要求常数时间获取最小值,存一个最小值(或下标)就行(官方题解是存最小栈，所有操作都常数时间)
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Stack: []int{}, MinIdx: 0}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.Stack) == 1 || (len(this.Stack) > 1 && val < this.Stack[this.MinIdx]) {
		this.MinIdx = len(this.Stack) - 1
	}
}

func (this *MinStack) Pop() {
	l := len(this.Stack)
	if this.MinIdx == l-1 {
		this.MinIdx = 0
		for i := 1; i < l-1; i++ {
			if this.Stack[i] < this.Stack[this.MinIdx] {
				this.MinIdx = i
			}
		}
	}
	this.Stack = this.Stack[:l-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Stack[this.MinIdx]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	s := []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}
	n := []int{0, -2, 1, -3, 0, 0, 0, 0}
	var stack MinStack
	for i := 0; i < len(s) && i < len(n); i++ {
		switch s[i] {
		case "MinStack":
			stack = Constructor()
			fmt.Println("MinStack")
		case "push":
			stack.Push(n[i])
			fmt.Println("push")
		case "getMin":
			fmt.Println("min:", stack.GetMin())
		case "pop":
			stack.Pop()
			fmt.Println("pop")
		case "top":
			fmt.Println("top", stack.Top())
		}
	}
}
