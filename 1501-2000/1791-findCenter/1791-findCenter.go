// https://leetcode-cn.com/problems/find-center-of-star-graph/
package main

import "fmt"

/**
 * 统计各个数字的个数，最多即为中心节点
 * 时间O(n)
 * 空间O(c), c<4
 */
func findCenter(edges [][]int) int {
	m := map[int]bool{}
	for i := 0; i < len(edges); i++ {
		for j := 0; j < len(edges[i]); j++ {
			if _, ok := m[edges[i][j]]; ok {
				return edges[i][j]
			}
			m[edges[i][j]] = true
		}
	}
	return 0
}

/**
 * 只需要判断前两组数据就行
 * 时间O(1)
 * 空间O(1)
 */
func findCenter2(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}

func main() {
	edgess := [][][]int{
		{{1, 2}, {2, 3}, {4, 2}},
		{{1, 2}, {5, 1}, {1, 3}, {1, 4}},
	}
	for _, edges := range edgess {
		fmt.Println(findCenter(edges), findCenter2(edges))
	}
}
