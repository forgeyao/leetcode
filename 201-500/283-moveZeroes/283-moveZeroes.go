// https://leetcode.cn/problems/move-zeroes/
package main

import "fmt"

/**
 * 双指针,一个负责找0，一个负责找非0
 * 时间 O(N)
 * 空间 O(1)
 */
func moveZeroes(nums []int) {
	l := len(nums)
	for i, j := 0, 0; i < l && j < l; {
		for ; i < l && nums[i] != 0; i++ {
		}
		for j = i + 1; j < l && nums[j] == 0; j++ {
		}

		if i >= l || j >= l {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}

/**
 * 官方答案。更简洁的双指针
 * 时间 O(N)
 * 空间 O(1)
 */
func moveZeroes2(nums []int) {
	i, j, l := 0, 0, len(nums)
	for j < l {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
func main() {
	numss := [][]int{
		{0},
		{1, 0},
		{0, 1},
		{1, 0, 1},
		{0, 1, 0, 3, 12},
		{0, 0, 0, 1, 2, 3},
	}
	for _, nums := range numss {
		tmp := nums
		fmt.Print(tmp)
		moveZeroes(nums)
		fmt.Print(" ", nums)
		moveZeroes2(tmp)
		fmt.Println(" ", tmp)
	}
}
