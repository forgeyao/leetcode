// https://leetcode-cn.com/problems/single-number/
package main

import "fmt"

func singleNumber(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret = ret ^ nums[i]
	}
	return ret
}

func main() {
	n := [][]int{{2, 2, 1}, {4, 1, 2, 1, 2}}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", singleNumber(n[i]))
	}
}
