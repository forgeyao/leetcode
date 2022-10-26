// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
package main

import "fmt"

/**
 * 双指针遍历
 * 时间 O(n)
 * 空间 O(1)
 */
func exchange(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]%2 == 0 && nums[j]%2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
			continue
		}
		if nums[i]%2 == 1 {
			i++
		}
		if nums[j]%2 == 0 {
			j--
		}
	}
	return nums
}

func main() {
	numss := [][]int{
		{1, 2, 3, 4},
		{11, 9, 3, 7, 16, 4, 2, 0},
	}
	results := [][]int{
		{1, 3, 2, 4},
		{11, 9, 3, 7, 16, 4, 2, 0},
	}

	for i := 0; i < len(numss) && i < len(results); i++ {
		fmt.Println(numss[i], "ret:", exchange(numss[i]), " result:", results[i])
	}
}
