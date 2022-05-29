// https://leetcode-cn.com/problems/convert-1d-array-into-2d-array/
package main

import "fmt"

// 一维数组转二维数组。常规思路，遍历一维数组，一个个复制到二维数组里
func construct2DArray(original []int, m int, n int) [][]int {
	l := len(original)
	if l != m*n {
		return [][]int{}
	}

	ret := make([][]int, m)
	for i := range ret {
		ret[i] = make([]int, n)
	}

	for i := 0; i < l; i++ {
		ret[i/n][i%n] = original[i]
	}
	return ret
}

// 方案一提交后，发现效率很低，做了改进，改成行复制
func construct2DArray2(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	ret := make([][]int, m)
	for i := 0; i < m; i++ {
		ret[i] = original[i*n : (i+1)*n]
	}
	return ret
}

func main() {
	original := [][]int{{1, 2, 3, 4}, {1, 2, 3}, {1, 2}}
	m := []int{2, 1, 1}
	n := []int{2, 3, 1}
	for i := 0; i < len(original) && i < len(m) && i < len(n); i++ {
		fmt.Println(original[i], " ret:", construct2DArray(original[i], m[i], n[i]), " ret2:", construct2DArray2(original[i], m[i], n[i]))
	}
}
