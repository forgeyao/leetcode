// https://leetcode-cn.com/problems/min-stack/
package main

import "fmt"

type MinStack struct {
	Stack []int
	Min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack MinStack
	stack.Stack = make([]int, 0)
	return stack
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.Stack) == 1 || (len(this.Stack) > 1 && val < this.Min) {
		this.Min = val
	}
}

func (this *MinStack) Pop() {
	l := len(this.Stack)
	if l == 0 {
		return
	}
	if this.Min == this.Stack[l-1] {
		if l == 1 {
			this.Min = 0
		} else {
			this.Min = this.Stack[0]

			for i := 0; i < l-1; i++ {
				if this.Stack[i] < this.Min {
					this.Min = this.Stack[i]
				}
			}
		}
	}
	this.Stack = this.Stack[:l-1]
}

func (this *MinStack) Top() int {
	var ret int
	if len(this.Stack) > 0 {
		ret = this.Stack[len(this.Stack)-1]
	}
	return ret
}

func (this *MinStack) GetMin() int {
	return this.Min
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
