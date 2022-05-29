// https://leetcode-cn.com/problems/pascals-triangle/
package main

import "fmt"

func generate(numRows int) [][]int {
	num := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		num[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				num[i][j] = 1
			} else {
				num[i][j] = num[i-1][j-1] + num[i-1][j]
			}
		}
	}
	return num
}

func main() {
	n := []int{5, 7}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", generate(n[i]))
	}
}
