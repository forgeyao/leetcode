// https://leetcode.cn/problems/binary-search/
package main

import "fmt"

/**
 * 二分查找，注意边界条件
 * 时间 O(logN)
 * 空间 O(1)
 */
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		//mid := (l+r)/2   // 相加存在越界问题
		mid := (r-l)/2 + l
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func main() {
	numss := [][]int{{-1, 0, 3, 5, 9, 12}, {-1, 0, 3, 5, 9, 12}, {1}, {1}}
	targets := []int{9, 2, 1, 0}
	answers := []int{4, -1, 0, -1}
	for i := 0; i < len(numss) && i < len(targets) && i < len(answers); i++ {
		fmt.Println(numss[i], targets[i], "ans:", answers[i], "ret:", search(numss[i], targets[i]))
	}
}
