// https://leetcode.cn/problems/sqrtx/
package main

import "fmt"

/**
 * 二分查找
 * 时间 O(logX)
 * 空间 O(1)
 */
func mySqrt(x int) int {
	left, right := 1, x
	for mid := left + (right-left)/2; left <= right; mid = left + (right-left)/2 {
		if mid == x/mid {
			return mid
		}
		if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 8, 9, 100, 1000}
	for i := 0; i < len(nums); i++ {
		fmt.Println("ret:", mySqrt(nums[i]))
	}
}
