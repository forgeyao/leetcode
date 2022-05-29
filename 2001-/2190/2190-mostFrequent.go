// https://leetcode-cn.com/problems/most-frequent-number-following-key-in-an-array/
package main

func mostFrequent(nums []int, key int) int {
	m := map[int]int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			m[nums[i+1]]++
		}
	}
	maxIdx := 0
	for k, v := range m {
		if v > m[maxIdx] {
			maxIdx = k
		}
	}
	return maxIdx
}
