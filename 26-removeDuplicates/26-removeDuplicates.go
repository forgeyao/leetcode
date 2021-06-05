// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[idx] == nums[i] {
			continue
		}
		idx++
		nums[idx] = nums[i]
	}
	nums = nums[:idx+1]
	return idx + 1
}

func main() {
	n := [][]int{{1, 1, 2}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", removeDuplicates(n[i]))
	}
}
