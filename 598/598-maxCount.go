// https://leetcode-cn.com/problems/range-addition-ii/
package main

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	minI, minJ := m, n
	for i := 0; i < len(ops); i++ {
		if len(ops[i]) >= 2 {
			minI = min(minI, ops[i][0])
			minJ = min(minJ, ops[i][1])
		}
	}
	return minI * minJ
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	m := []int{3, 4}
	n := []int{3, 4}
	o := [][][]int{{{2, 2}, {3, 3}}, {{1, 3}, {4, 2}, {2, 2}}}
	for i := 0; i < len(m) && i < len(n) && i < len(o); i++ {
		fmt.Println("ret:", maxCount(m[i], n[i], o[i]))
	}
}
