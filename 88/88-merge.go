// https://leetcode-cn.com/problems/merge-sorted-array/
package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, k := m-1, n-1, len(nums1)-1; k >= 0; k-- {
		if j < 0 {
			break
		}

		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}

func main() {
	nums1 := [][]int{{1, 2, 3, 0, 0, 0}, {1}}
	nums2 := [][]int{{2, 5, 6}, {}}
	m := []int{3, 1}
	n := []int{3, 0}
	for i := 0; i < len(nums1) && i < len(nums2) && i < len(m) && i < len(n); i++ {
		merge(nums1[i], m[i], nums2[i], n[i])
		fmt.Println("ret:", nums1[i])
	}
}
