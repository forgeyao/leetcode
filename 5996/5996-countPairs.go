// https://leetcode-cn.com/problems/count-equal-and-divisible-pairs-in-an-array/
package main

import "fmt"

func countPairs(nums []int, k int) int {
	ret := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				ret++
			}
		}
	}
	return ret
}

func main() {
	nums := [][]int{
		{3, 1, 2, 2, 2, 1, 3},
		{1, 2, 3, 4},
	}
	k := []int{2, 1}
	for i := 0; i < len(nums) && i < len(k); i++ {
		fmt.Println(countPairs(nums[i], k[i]))
	}
}
