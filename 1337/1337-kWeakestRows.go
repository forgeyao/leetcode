// https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
package main

import "fmt"

func kWeakestRows(mat [][]int, k int) []int {
	ret := []int{}
	row := len(mat)
	col := len(mat[0])
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if mat[j][i] == 0 && (i == 0 || (i > 0 && mat[j][i-1] == 1)) {
				ret = append(ret, j)
				if len(ret) >= k {
					return ret
				}
			}
		}

		if i == col-1 {
			for j := 0; j < row; j++ {
				if mat[j][i] == 1 {
					ret = append(ret, j)
					if len(ret) >= k {
						return ret
					}
				}
			}
		}
	}
	return ret
}

func main() {
	mat1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	mat2 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	mat3 := [][]int{
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
	}
	mat4 := [][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
	}
	k := []int{3, 2, 1, 3}

	fmt.Println("ret:", kWeakestRows(mat1, k[0]))
	fmt.Println("ret:", kWeakestRows(mat2, k[1]))
	fmt.Println("ret:", kWeakestRows(mat3, k[2]))
	fmt.Println("ret:", kWeakestRows(mat4, k[3]))
}
