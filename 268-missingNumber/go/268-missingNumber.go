// https://leetcode-cn.com/problems/missing-number/
package main

import "fmt"

/**
 * 用求和公式算出缺少的值
 */
func missingNumber(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return len(nums)*(len(nums)+1)/2 - sum
}

/**
 * 优化版
 */
func missingNumber2(nums []int) int {
	miss := len(nums)
	for i := 0; i < len(nums); i++ {
		miss += i - nums[i]
	}
	return miss
}

func main() {
	nums := [][]int{{3, 0, 1}, {9, 6, 4, 2, 3, 5, 7, 0, 1}, {0}, {0, 1}}
	for _, num := range nums {
		fmt.Println("ret:", missingNumber(num), missingNumber2(num))
	}
}
