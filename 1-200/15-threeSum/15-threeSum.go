// https://leetcode.cn/problems/3sum/
package main

import (
	"fmt"
	"sort"
)

/**
 * 排序+3层遍历。结果超时
 * 时间: O(N*N*N)
 * 空间: O(1)
 */
func threeSum(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	l := len(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < l; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ret
}

/**
 * 排序+3层遍历。细节有优化
 * 时间: O(N*N)
 * 空间: O(1)
 */

func threeSum2(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	l := len(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		thirdIdx := l - 1
		for j := i + 1; j < l-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := thirdIdx; k > j; k-- {
				//fmt.Println(i, j, k, "num:", nums[i], nums[j], nums[k])
				if k < l-1 && nums[k] == nums[k+1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
					thirdIdx = k - 1
					break
				} else if nums[i]+nums[j]+nums[k] < 0 {
					break
				}
			}
		}
	}
	return ret
}

func main() {
	nums := [][]int{
		/*{-1, 0, 1, 2, -1, -4},
		{0, 0, 0},
		{0, 0, 0, 0},
		{-2, 0, 0, 2, 2},
		{3, 0, -2, -1, 1, 2},
		{-2, 0, 1, 1, 2},*/
		{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
	}
	results := [][][]int{
		/*{{-1, 0, 1}, {-1, -1, 2}},
		{{0, 0, 0}},
		{{0, 0, 0}},
		{{-2, 0, 2}},
		{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		{{-2, 0, 2}, {-2, 1, 1}},*/
		{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
	}
	for i := 0; i < len(nums) && i < len(results); i++ {
		fmt.Println(threeSum2(nums[i]), "result:", results[i])
	}
}
