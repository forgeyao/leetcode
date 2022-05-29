// https://leetcode-cn.com/problems/next-greater-element-i/
package main

import "fmt"

/**
 * 将无序数组存到 map 来加快搜索
 */
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	m := make(map[int]int, l2)
	for i := 0; i < l2; i++ {
		m[nums2[i]] = i
	}

	ret := make([]int, l1)
	for i := 0; i < l1; i++ {
		idx, _ := m[nums1[i]]
		ret[i] = -1
		for j := idx + 1; j < l2; j++ {
			if nums2[j] > nums1[i] {
				ret[i] = nums2[j]
				break
			}
		}
	}
	return ret
}

/**
 * 单向栈 + map
 */
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	l1, l2 := len(nums1), len(nums2)
	m := make(map[int]int, l2)
	stack := []int{}
	for i := l2 - 1; i >= 0; i-- {
		// 栈不为空 且栈顶较小 较小的均出栈
		idx := len(stack) - 1
		for ; idx >= 0 && stack[idx] < nums2[i]; idx-- {
		}
		stack = stack[0 : idx+1]

		// 此时栈要么为空 要么里面的元素均大于最新元素
		if len(stack) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i]) // 最新元素入栈
	}

	ret := make([]int, l1)
	for i := 0; i < l1; i++ {
		ret[i] = m[nums1[i]]
	}
	return ret
}

func main() {
	nums1 := [][]int{{4, 1, 2}, {2, 4}, {1, 3, 5, 2, 4}}
	nums2 := [][]int{{1, 3, 4, 2}, {1, 2, 3, 4}, {6, 5, 4, 3, 2, 1, 7}}
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		fmt.Println("ret:", nextGreaterElement(nums1[i], nums2[i]), nextGreaterElement2(nums1[i], nums2[i]))
	}
}
