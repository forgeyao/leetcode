// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
package main

import "fmt"

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

func main() {
	n := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	fmt.Println("ret:", luckyNumbers(n))
	n2 := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	fmt.Println("ret:", luckyNumbers(n2))
	n3 := [][]int{{7, 8}, {1, 2}}
	fmt.Println("ret:", luckyNumbers(n3))
}
