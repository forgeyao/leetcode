// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
package main

import "fmt"

type CQueue struct {
	a []int
	b []int
}

/**
 * 一个作为入栈,一个作为出栈
 * 时间: 入栈O(1),出栈平均O(1)
 * 空间: O(N), 栈深度最大N
 */
func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.a = append(this.a, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.b) == 0 {
		if len(this.a) == 0 {
			return -1
		}

		for len(this.a) > 0 {
			la := len(this.a)
			this.b = append(this.b, this.a[la-1])
			this.a = this.a[:la-1]
		}
	}
	ret := this.b[len(this.b)-1]
	this.b = this.b[:len(this.b)-1]
	return ret
}

func main() {
	fmt.Println("vim-go")
}
