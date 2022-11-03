// https://leetcode.cn/problems/majority-element/
package main

import "fmt"

/**
 * hash
 * 时间 O(n)
 * 空间 O(n)
 */
func majorityElement(nums []int) int {
	m := make(map[int]int, len(nums)/2+1)
	for _, val := range nums {
		if _, ok := m[val]; ok {
			m[val] += 1
		} else {
			m[val] = 1
		}
		if m[val] > len(nums)/2 {
			return val
		}
	}
	return 0
}

/**
 * Boyer-Moore 投票算法
 * 时间 O(n)
 * 空间 O(1)
 */
func majorityElement2(nums []int) int {
	majority, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if majority == nums[i] {
			count++
		} else {
			count--
			if count == 0 {
				majority = nums[i]
				count = 1
			}
		}
	}
	return majority
}

func main() {
	n := [][]int{{3, 2, 3}, {2, 2, 1, 1, 1, 2, 2}, {1, 1, 1, 2, 2, 2, 2}}
	for _, val := range n {
		fmt.Println("ret:", majorityElement(val), majorityElement2(val))
	}
}
