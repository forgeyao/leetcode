// https://leetcode.cn/problems/search-in-rotated-sorted-array/
package main

import "fmt"

/**
 * 二分后，两个子数组必有一个是有序的
 * 如果在有序数组范围内，则二分查找
 * 如果不在有序数组，继续二分
 * 迭代前两步

 * 时间: O(log(N)
 * 空间: O(1)
 */
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	numss := [][]int{{4, 5, 6, 7, 0, 1, 2}, {4, 5, 6, 7, 0, 1, 2}, {1}, {5, 1, 3}, {3, 1}}
	targets := []int{0, 3, 0, 5, 1}
	answers := []int{4, -1, -1, 0, 1}
	for i := 0; i < len(numss) && i < len(targets); i++ {
		fmt.Println(numss[i], targets[i], "ans:", answers[i], "ret:", search(numss[i], targets[i]))
	}
}
