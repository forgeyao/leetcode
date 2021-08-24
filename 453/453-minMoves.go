// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
package main

import "fmt"

func minMoves(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	count := 0
	for _, v := range nums {
		count += v - min
	}
	return count
}

func main() {
	n := [][]int{{1, 2, 3}, {0, 2, 4}, {1, 4, 4}}
	for _, v := range n {
		fmt.Println("ret:", minMoves(v))
	}
}
