// https://leetcode-cn.com/problems/optimal-division/
package main

import (
	"fmt"
	"strconv"
)

/**
 * 思路：要想结果最大，就让分母最小。要让分母最小，分母(本身也是分数)的分母最大，依次循环
 * 递归
 */
func optimalDivision(nums []int) string {
	var f func(nums []int, b bool) string
	f = func(nums []int, b bool) string {
		if len(nums) <= 2 {
			s := ""
			for i := 0; i < len(nums); i++ {
				if i%2 == 1 {
					s += "/" + strconv.Itoa(nums[1])
				} else {
					s += strconv.Itoa(nums[i])
				}
			}
			return s
		}
		if b {
			return strconv.Itoa(nums[0]) + "/(" + f(nums[1:], false) + ")"
		}
		return strconv.Itoa(nums[0]) + "/" + f(nums[1:], false)
	}

	return f(nums, true)
}

/**
 * 写完递归测试发现只要给分母家括号就行 😂
 * 时间 O(n)
 * 空间 O(1)
 */
func optimalDivision2(nums []int) string {
	s := strconv.Itoa(nums[0])
	l := len(nums)
	for i := 1; i < l; i++ {
		if l >= 3 && i == 1 {
			s += "/(" + strconv.Itoa(nums[i])
		} else {
			s += "/" + strconv.Itoa(nums[i])
		}
	}
	if l >= 3 {
		s += ")"
	}
	return s
}

func main() {
	nums := [][]int{
		{1000, 100, 10, 2},
		{6, 2, 3, 4, 5},
	}
	for _, num := range nums {
		fmt.Println(optimalDivision(num), optimalDivision2(num))
	}
}
