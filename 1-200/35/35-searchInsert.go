// https://leetcode-cn.com/problems/search-insert-position/
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := left + (right-left)/2
	for left <= mid && mid <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if left > right {
			return left
		}
		mid = left + (right-left)/2
	}
	return 0
}

func main() {
	nums := [][]int{{1, 3, 5, 6}, {1, 3, 5, 6}, {1, 3, 5, 6}, {1, 3, 5, 6}, {1, 3, 5, 6}}
	n := []int{5, 2, 7, 0, 4}
	for i := 0; i < len(nums) && i < len(n); i++ {
		fmt.Println("ret:", searchInsert(nums[i], n[i]))
	}
}
