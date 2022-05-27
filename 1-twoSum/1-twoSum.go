// https://leetcode.cn/problems/two-sum/submissions/
package main

import (
	"fmt"
	"sort"
)

/**
 * 没直接用蛮力法，尝试用排序来解
 *
 * 时间 O(N*logN), 排序时间复杂度
 * 空间 O(N)
 */
func twoSum(nums []int, target int) []int {
	m := map[int][]int{}
	for i, v := range nums {
		if val, ok := m[v]; ok {
			val = append(val, i)
			m[v] = val
		} else {
			m[v] = []int{i}
		}
	}
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			if nums[i] == nums[j] {
				return []int{m[nums[i]][0], m[nums[i]][1]}
			}
			return []int{m[nums[i]][0], m[nums[j]][0]}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

/**
 * 参考官方答案思路，优化 twoSum。无序排序
 *
 * 时间 O(N)
 * 空间 O(N)
 */
func twoSum2(nums []int, target int) []int {
	m := map[int][]int{}
	for i, v := range nums {
		if val, ok := m[v]; ok {
			val = append(val, i)
			m[v] = val
		} else {
			m[v] = []int{i}
		}
	}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		v, ok := m[diff]
		if ok {
			if len(v) == 1 && v[0] != i {
				return []int{i, v[0]}
			}
			if len(v) > 1 {
				return []int{i, v[1]}
			}
		}
	}
	return []int{}
}

/**
 * 官方答案
 *
 * 时间 O(N)
 * 空间 O(N)
 */
func twoSum3(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := m[diff]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func main() {
	numss := [][]int{{2, 7, 11, 15}, {3, 2, 4}, {3, 3}}
	targets := []int{9, 6, 6}
	for i := 0; i < len(numss) && i < len(targets); i++ {
		// twoSum 有排序，要后执行
		fmt.Println("ret:", twoSum2(numss[i], targets[i]), twoSum3(numss[i], targets[i]), twoSum(numss[i], targets[i]))
	}
}
