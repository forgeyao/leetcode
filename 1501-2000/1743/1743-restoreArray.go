// https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs/
package main

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
	l := len(adjacentPairs) + 1
	m := make(map[int][]int, l) // 相邻映射关系
	for _, v := range adjacentPairs {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}

	ret := make([]int, l)
	for k, v := range m {
		if len(v) == 1 { // 头或者尾
			ret[0] = k
			ret[1] = m[k][0]
			break
		}
	}
	for i := 2; i < l; i++ {
		val := m[ret[i-1]]
		if len(val) == 2 && val[0] == ret[i-2] {
			ret[i] = val[1]
		} else {
			ret[i] = val[0]
		}
	}
	return ret
}

func main() {
	n := [][][]int{{{2, 1}, {3, 4}, {3, 2}}, {{4, -2}, {1, 4}, {-3, 1}}, {{100000, -100000}}, {{-3, -9}, {-5, 3}, {2, -9}, {6, -3}, {6, 1}, {5, 3}, {8, 5}, {-5, 1}, {7, 2}}}
	for _, v := range n {
		fmt.Println("ret:", restoreArray(v))
	}
}
