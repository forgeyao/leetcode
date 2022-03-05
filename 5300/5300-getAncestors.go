//https://leetcode-cn.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/
package main

import (
	"fmt"
	"sort"
)

func getAncestors(n int, edges [][]int) [][]int {
	m := map[int][]int{}
	for _, edge := range edges {
		if v, ok := m[edge[1]]; ok {
			v = append(v, edge[0])
			m[edge[1]] = v
		} else {
			arr := []int{edge[0]}
			m[edge[1]] = arr
		}
	}
	ret := [][]int{}
	for i := 0; i < n; i++ {
		r := []int{}
		if v, ok := m[i]; ok {
			ret = append(ret, uniqueSort(v, m))
		} else {
			ret = append(ret, r)
		}
	}
	return ret
}

func uniqueSort(num []int, mm map[int][]int) []int {
	m := map[int]int{}
	ret := []int{}
	arr := num
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
		if m[arr[i]] == 1 {
			ret = append(ret, arr[i])
			arr = append(arr, mm[arr[i]]...)
		}
	}
	sort.Ints(ret)
	return ret
}

func main() {
	params := []struct {
		n     int
		edges [][]int
	}{
		{n: 8, edges: [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}},
		{n: 5, edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}
	for _, item := range params {
		fmt.Println(getAncestors(item.n, item.edges))
	}
}
