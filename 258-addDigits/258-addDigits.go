// https://leetcode-cn.com/problems/add-digits/
package main

import (
	"fmt"
)

/**
 * 思路：除法、取余计算
 * 时间 O(log num) num 为给定整数
 * 空间 O(1)
 */
func AddDigits(num int) int {
	for {
		sum := 0
		for num != 0 {
			sum += num % 10
			num /= 10
		}
		if sum/10 == 0 {
			return sum
		}
		num = sum
	}
}

/**
 * 参考官方答案
 * 时间 O(1)
 * 空间 O(1)
 */
func AddDigits2(num int) int {
	return (num-1)%9 + 1
}

func main() {
	nums := []int{38, 0}
	results := []int{2, 0}
	for i := 0; i < len(nums) && i < len(results); i++ {
		ret := AddDigits(nums[i])
		if ret != results[i] {
			fmt.Println("num:", nums[i], "ret:", ret, "result:", results[i])
		}
		fmt.Println(AddDigits2(nums[i]))
	}
}
