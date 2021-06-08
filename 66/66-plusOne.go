// https://leetcode-cn.com/problems/plus-one/
package main

import "fmt"

func plusOne(digits []int) []int {
	l := len(digits)
	if digits[l-1] != 9 {
		digits[l-1] += 1
		return digits
	}

	digits[l-1] = 0
	flag := 1
	for i := len(digits) - 2; i >= 0; i-- {
		if digits[i]+flag < 10 {
			digits[i] += flag
			return digits
		}
		digits[i] = 0
		flag = 1
	}
	if flag == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func main() {
	nums := [][]int{{1, 2, 3}, {4, 3, 2, 1}, {1, 2, 9}, {9, 9, 9}}
	for i := 0; i < len(nums); i++ {
		fmt.Println("ret:", plusOne(nums[i]))
	}
}
