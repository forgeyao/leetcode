// https://leetcode-cn.com/problems/contains-duplicate/
package main

import "fmt"

// 哈希表
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

func main() {
	n := [][]int{{1, 2, 3, 1}, {1, 2, 3, 4}, {1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}
	for _, v := range n {
		fmt.Println("ret:", containsDuplicate(v))
	}
}
