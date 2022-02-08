// https://leetcode-cn.com/problems/path-with-maximum-gold/
package main

import "fmt"

/**
 * 有想到递归，但没想好怎么实现
 * 下面看参考答案的
 */
var dirs = []struct{ x, y int }{
	{-1, 0}, // 上
	{1, 0},  // 下
	{0, -1}, // 左
	{0, 1},  // 右
}

func getMaximumGold(grid [][]int) (ans int) {
	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		gold += grid[x][y]
		if gold > ans {
			ans = gold
		}

		rec := grid[x][y]
		grid[x][y] = 0 // 标记已开采过
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[nx]) && grid[nx][ny] > 0 {
				dfs(nx, ny, gold)
			}
		}
		grid[x][y] = rec // 恢复
	}

	for i, row := range grid {
		for j, gold := range row {
			if gold > 0 {
				dfs(i, j, 0)
			}
		}
	}
	return
}

func main() {
	grid := [][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}
	grid2 := [][]int{
		{1, 0, 7},
		{2, 0, 6},
		{3, 4, 5},
		{0, 3, 0},
		{9, 0, 20},
	}
	fmt.Println(getMaximumGold(grid), getMaximumGold(grid2))
}
