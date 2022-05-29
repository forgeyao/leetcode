// https://leetcode-cn.com/problems/number-of-enclaves/
package main

import "fmt"

/**
 * 思路: 用一个同等大小二维数组保存 grid 为 1 时，能离开边界状态。初始化二位数组，四边为1 状态都为 true
 * 遍历 grid，当为1，并且上下左右有 flag 为true时，该flag 为 true，递归上下左右进一步扩散 flag
 * 最终统计 grid 为1，flag 为false 的个数
 * 看官方答案，才发现这方法就是深度优先搜索(DFS)
 * 时间O(mn)，其中 m 和 n 分别是网格grid的行数和列数
 * 空间O(mn)
 */
func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	flag := make([][]bool, m) // flag 表示 grid 为 1 时，能离开边界

	for i := 0; i < m; i++ {
		flag[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			flag[i][j] = false
			if grid[i][j] == 1 && (i == 0 || i == m-1 || j == 0 || j == n-1) { // 边界为1的，flag 均为 true
				flag[i][j] = true
			}
		}
	}
	/*for i := 0; i < m; i++ {
		fmt.Println("init:", flag[i])
	}*/

	var visit func(i, j int)
	visit = func(i, j int) {
		// 边界、grid 为 0、flag 为 true 均退出
		if (i == 0 || i == m-1 || j == 0 || j == n-1) || grid[i][j] == 0 || flag[i][j] {
			return
		}
		// 上下左右有 flag 为 true，则 flag[i][j] 为 true(这里的 grid[i][j] 都为 1)
		// 继续向四周扩散 flag
		if flag[i-1][j] || flag[i+1][j] || flag[i][j-1] || flag[i][j+1] {
			flag[i][j] = true
			visit(i-1, j) // 上
			visit(i+1, j) // 下
			visit(i, j-1) // 左
			visit(i, j+1) // 右
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			visit(i, j)
		}
	}

	/*for i := 0; i < m; i++ {
		fmt.Println("end:", flag[i])
	}*/
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !flag[i][j] { // grid 为 1, flag 为 false 即为所求
				count++
			}
		}
	}
	return count
}

func main() {
	//grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	//grid := [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	grid := [][]int{{0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {1, 1, 0, 0, 0, 1, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {0, 1, 1, 0, 0, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1, 0, 0, 1, 0}, {0, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, {0, 0, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}
	fmt.Println(numEnclaves(grid))
}
