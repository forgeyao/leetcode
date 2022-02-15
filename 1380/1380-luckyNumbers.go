// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
package main

import "fmt"

/**
 * 时间O(m*n),m 是行数,n 是列数
 * 空间O(m), 缓存行里最小值
 */
func luckyNumbers(matrix [][]int) []int {
	col := len(matrix)
	row := len(matrix[0])
	lucky := make(map[int]bool, row)
	for i := 0; i < col; i++ {
		min := matrix[i][0]
		for j := 1; j < row; j++ {
			if min > matrix[i][j] {
				min = matrix[i][j]
			}
		}
		lucky[min] = true
	}
	ret := []int{}
	for i := 0; i < row; i++ {
		max := matrix[0][i]
		for j := 1; j < col; j++ {
			if max < matrix[j][i] {
				max = matrix[j][i]
			}
		}
		if _, ok := lucky[max]; ok {
			ret = append(ret, max)
		}
	}
	return ret
}

/**
 * 优化掉空间
 * 时间O(m*n),m 是行数,n 是列数
 * 空间O(1)
 */
func luckyNumbers2(matrix [][]int) []int {
	ret := []int{}
	for i := 0; i < len(matrix); i++ {
		minJ := 0
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < matrix[i][minJ] {
				minJ = j
			}
		}
		b := false
		for k := 0; k < len(matrix); k++ {
			if matrix[i][minJ] < matrix[k][minJ] {
				b = false
				break
			}
			b = true
		}
		if b {
			ret = append(ret, matrix[i][minJ])
		}
	}
	return ret
}

func main() {
	ns := [][][]int{
		{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
		{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}},
		{{7, 8}, {1, 2}},
	}
	for _, n := range ns {
		fmt.Println("ret:", luckyNumbers(n), luckyNumbers2(n))
	}
}
