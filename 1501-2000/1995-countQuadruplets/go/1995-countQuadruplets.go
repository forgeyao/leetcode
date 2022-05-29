// https://leetcode-cn.com/problems/count-special-quadruplets/
package main

import "fmt"

/**
 * 将求解4元组转换为求3元组，逐步分解问题
 * 使用递归
 */
func countQuadruplets(nums []int) int {
	count := 0
	// 遍历 nums，找所有和为 nums[i] 的3元组
	for i := len(nums) - 1; i >= 3; i-- {
		count += countNum(nums[:i], 3, nums[i])
	}
	return count
}

// 在 nums 找 n 个元素, 和为 sum
func countNum(nums []int, n int, sum int) int {
	l := len(nums)
	if l == 0 || l < n {
		return 0
	}
	count := 0
	if n == 1 {
		for _, v := range nums {
			if v == sum {
				count++
			}
		}
		return count
	}

	for i := l - 1; i >= n-1; i-- {
		if nums[i] >= sum {
			continue
		}
		// n 元组分解为 n-1 元组
		count += countNum(nums[:i], n-1, sum-nums[i])
	}
	return count
}

func main() {
	nums := [][]int{{1, 2, 3, 6}, {3, 3, 6, 4, 5}, {1, 1, 1, 3, 5}}
	for _, num := range nums {
		fmt.Println(num, " cout: ", countQuadruplets(num))
	}
}
