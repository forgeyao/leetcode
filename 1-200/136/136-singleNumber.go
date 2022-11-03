// https://leetcode.cn/problems/single-number/
package main

import "fmt"

/**
 * 异或运算
 * 时间 O(n)
 * 空间 O(1)
 */
func singleNumber(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}

func main() {
	n := [][]int{{2, 2, 1}, {4, 1, 2, 1, 2}}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", singleNumber(n[i]))
	}
}
