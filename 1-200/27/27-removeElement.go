// https://leetcode-cn.com/problems/remove-element/
package main

import "fmt"

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[j] == val {
			j--
			continue
		}
		if nums[i] == val {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			j--
		}
		i++
	}
	return i
}

func main() {
	n := [][]int{{3, 2, 2, 3}, {0, 1, 2, 2, 3, 0, 4, 2}, {1, 2, 3, 4}}
	v := []int{3, 2, 5}
	for i := 0; i < len(n) && i < len(v); i++ {
		fmt.Println("ret:", removeElement(n[i], v[i]))
	}
}
