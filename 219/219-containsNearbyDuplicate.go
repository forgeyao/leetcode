// https://leetcode-cn.com/problems/contains-duplicate-ii/
package main

import "fmt"

// 缓存
// 时间：O(n)  空间：O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if arr, ok := m[nums[i]]; ok {
			arr = append(arr, i)
			m[nums[i]] = arr
		} else {
			newArr := []int{i}
			m[nums[i]] = newArr
		}
	}
	for _, v := range m {
		//fmt.Println("k,v:", key, v)
		for i := 1; i < len(v); i++ {
			if v[i]-v[i-1] <= k {
				return true
			}
		}
	}
	return false
}

// 遍历
// 时间：O(n*min(n,k))  空间：O(1)
func containsNearbyDuplicate2(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}

// 缓存优化版
// 时间：O(n)  空间：O(n)
func containsNearbyDuplicate3(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if index, ok := m[nums[i]]; ok {
			if i-index <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}

func main() {
	n := [][]int{{1, 2, 3, 1}, {1, 0, 1, 1}, {1, 2, 3, 1, 2, 3}}
	k := []int{3, 1, 2}
	for i := 0; i < len(n) && i < len(k); i++ {
		fmt.Println("ret:", containsNearbyDuplicate(n[i], k[i]), containsNearbyDuplicate2(n[i], k[i]), containsNearbyDuplicate3(n[i], k[i]))
	}
}
