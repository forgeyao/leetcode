// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	sum := make([]int, len(nums)) // 直接用 nums 可以省掉 sum
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i] // 前缀和
	}
	min, max, ret := sum[0], sum[0], sum[0]
	for i := 1; i < len(sum); i++ {
		if sum[i] <= min {
			min = sum[i]
			max = sum[i]
			if ret < nums[i] { // nums[i] = sum[i] - sum[i-1]
				ret = nums[i]
			}
		}
		if sum[i] > max {
			max = sum[i]
			if ret < max-min {
				ret = max - min
			}
			if ret < max {
				ret = max
			}
		}
	}
	return ret
}

// 优化版
func maxSubArray2(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i]+nums[i-1])
		ret = max(ret, nums[i])
	}
	return ret
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := [][]int{{-2}, {-2, 1}, {-1, 0}, {-2, -1}, {2, 1}, {1, 2}, {5, -1, -5, 1, 3}, {-2, 1, -3, 4, -1, 2, 1, -5, 4}}
	for _, v := range n {
		fmt.Println("ret:", maxSubArray(v), maxSubArray2(v))
	}
}
