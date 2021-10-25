// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
package main

import "fmt"

/**
 * 注意：本方法是错的！这里是为了记录解题思路
 *
 * 本来想通过二分查找加快查询
 * 但有漏洞，测试用例5、6 通不过
 * 核心问题是通过左上角到右下角的常规遍历，不能很好排除数据
 */

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l1, l2 := 0, 0
	r1, r2 := m-1, n-1
	for r1 >= l1 && r2 >= l2 {
		m1, m2 := (r1+l1)/2, (r2+l2)/2
		if target == matrix[m1][m2] {
			return true
		} else if target > matrix[m1][m2] {
			if m1 == r1 {
				if m2 == r2 { // 矩阵 1 x 1
					return false
				} else { // 矩阵 1 x N
					l2 = m2 + 1
				}
				continue
			}
			if m2 == r2 && m1 < r1 { // 矩阵 N x 1
				l1 = m1 + 1
				continue
			}

			// 矩阵 M x N
			if target == matrix[m1+1][l2] || target == matrix[l1][m2+1] {
				return true
			} else if target > matrix[m1+1][l2] && target > matrix[l1][m2+1] {
				l1 = m1 + 1
				l2 = m2 + 1
			} else {
				return false
			}
		} else {
			if m1 == l1 {
				if m2 == l2 {
					return false
				} else {
					r2 = m2 - 1
				}
				continue
			}
			if m2 == l2 && m1 > l1 {
				r1 = m1 - 1
				continue
			}
			if target == matrix[m1-1][r2] || target == matrix[r1][m2-1] {
				return true
			} else if target < matrix[m1-1][r2] && target < matrix[r1][m2-1] {
				r1 = m1 - 1
				r2 = m2 - 1
			} else {
				return false
			}
		}
	}
	return false
}

/**
 * 与常规从左上角或右下角遍历不同，这里从右上角开始遍历
 * 好处是如果小与 target, 则可以排除一行
 * 如果大于 target , 则可以排除一列
 * 持续迭代就能找到结果
 */
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {
	matrix := [][][]int{
		// test 1
		{
			{1},
		},
		// test 2
		{
			{1, 2},
		},
		// test 3
		{
			{1},
			{2},
		},
		// test 4
		{
			{1, 2},
			{3, 4},
		},

		// test 5
		{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		},

		// test 6
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
	}
	target := []int{5, 5, 5, 5, 15, 15}
	for i := 0; i < len(matrix) && i < len(target); i++ {
		fmt.Println("ret:", searchMatrix(matrix[i], target[i]), searchMatrix2(matrix[i], target[i]))
	}
}
