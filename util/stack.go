package util

type Stack struct {
	Val []int
}

func Constructor() Stack {
	return Stack{[]int{}}
}

func (this *Stack) Push(n int) {
	this.Val = append(this.Val, n)
}

func (this *Stack) Peek() int {
	size := this.Size()
	if size > 0 {
		val := this.Val[size-1]
		return val
	}
	return 0
}

func (this *Stack) Pop() int {
	size := this.Size()
	if size > 0 {
		val := this.Val[size-1]
		this.Val = this.Val[:size-1]
		return val
	}
	return 0
}

func (this *Stack) Size() int {
	return len(this.Val)
}

func (this *Stack) Empty() bool {
	return len(this.Val) == 0
}
