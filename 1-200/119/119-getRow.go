// https://leetcode-cn.com/problems/pascals-triangle-ii/
package main

import "fmt"

func getRow(rowIndex int) []int {
	// 构造完整的杨辉三角形. 二维数组
	/*num := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		num[i] = make([]int, i+1)
		num[i][0] = 1
		num[i][i] = 1
		for j := 1; j < i; j++ {
			num[i][j] = num[i-1][j-1] + num[i-1][j]
		}
	}
	return num[rowIndex]*/

	// 节省空间版。一维数组
	num := make([]int, rowIndex+1)
	num[0] = 1
	for i := 1; i <= rowIndex; i++ {
		num[i] = 1
		pre := num[0]
		for j := 1; j < i; j++ {
			cur := num[j]
			num[j] = pre + cur
			pre = cur
		}
	}
	return num
}

func main() {
	n := []int{3, 4, 5}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", getRow(n[i]))
	}
}
