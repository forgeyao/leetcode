// https://leetcode.cn/problems/number-of-islands/
package main

import "fmt"

/**
 * 从左到右、上到下遍历二维数组
 * 如果当前值不是'1',跳过
 * 如果为'1'，判断左和上是否为'0'，右和下是否没访问过，
 * 都是，则到找到一个岛的左上角，并相邻点逐步找到岛的全部
 *
 * 访问过的'1'会通过自增来标记已访问
 * 本质上就是广度优先搜索
 *
 * 时间 O(M*N), M 和 N 分别为行数和列数
 * 空间 O(min(M,N))
 */
func numIslands(grid [][]byte) int {
	num := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' || grid[i][j] > '1' {
				continue
			}
			if (j == 0 || j > 0 && grid[i][j-1] == '0') &&
				(i == 0 || i > 0 && grid[i-1][j] == '0') &&
				(j == n-1 || j < n-1 && grid[i][j+1] <= '1') &&
				(i == m-1 || i < m-1 && grid[i+1][j] <= '1') {
				num++
				grid[i][j]++ // 大于'1'标记为已访问过
				visitIsland(i, j, grid)
			}
		}
	}

	return num
}

type Pos struct {
	i, j int
}

func visitIsland(ii, jj int, grid [][]byte) {
	p := []Pos{{ii, jj}} // 模拟栈
	m, n := len(grid), len(grid[0])
	for len(p) > 0 {
		i, j := p[0].i, p[0].j
		p = p[1:]
		if j > 0 && grid[i][j-1] == '1' {
			grid[i][j-1]++
			p = append(p, Pos{i, j - 1})
		}
		if i > 0 && grid[i-1][j] == '1' {
			grid[i-1][j]++
			p = append(p, Pos{i - 1, j})
		}
		if i < m-1 && grid[i+1][j] == '1' {
			grid[i+1][j]++
			p = append(p, Pos{i + 1, j})
		}
		if j < n-1 && grid[i][j+1] == '1' {
			grid[i][j+1]++
			p = append(p, Pos{i, j + 1})
		}
	}
}

/**
 * 第二次实现,总体思路还是一样
 * 时间 O(M*N), M 和 N 分别为行数和列数
 * 空间 O(min(M,N))
 */
func numIslands2(grid [][]byte) int {
	visited := [][]bool{} // 标记已访问过
	for i := 0; i < len(grid); i++ {
		visited = append(visited, []bool{})
		for j := 0; j < len(grid[0]); j++ {
			visited[i] = append(visited[i], false)
		}
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visited[i][j] == true {
				continue
			}
			if grid[i][j] == '1' {
				count++
				toVisit := []Pos{Pos{i, j}}
				visit(grid, visited, toVisit)
			} else {
				visited[i][j] = true
			}
		}
	}
	return count
}

// 广度优先搜索，找出一个完整的岛
func visit(grid [][]byte, visited [][]bool, toVisit []Pos) {
	for i := 0; i < len(toVisit); i++ {
		pos := toVisit[i]
		if pos.i < 0 || pos.i >= len(grid) || pos.j < 0 || pos.j >= len(grid[0]) {
			continue
		}
		if visited[pos.i][pos.j] == true {
			continue
		}

		visited[pos.i][pos.j] = true
		if grid[pos.i][pos.j] == '0' {
			continue
		}
		toVisit = append(toVisit, Pos{pos.i - 1, pos.j})
		toVisit = append(toVisit, Pos{pos.i + 1, pos.j})
		toVisit = append(toVisit, Pos{pos.i, pos.j - 1})
		toVisit = append(toVisit, Pos{pos.i, pos.j + 1})
	}
}

/**
 * 深度有限搜索实现
 * 时间 O(M*N), M 和 N 分别为行数和列数
 * 空间 O(M*N)), 最坏情况，整个网络是一个岛，搜索深度达到M*N
 */
func numIslands3(grid [][]byte) int {
	visited := [][]bool{} // 标记已访问过
	for i := 0; i < len(grid); i++ {
		visited = append(visited, []bool{})
		for j := 0; j < len(grid[0]); j++ {
			visited[i] = append(visited[i], false)
		}
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visited[i][j] == true {
				continue
			}
			if grid[i][j] == '1' {
				count++
				dfs(grid, visited, i, j)
			} else {
				visited[i][j] = true
			}
		}
	}
	return count
}

func dfs(grid [][]byte, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if visited[i][j] == true {
		return
	}
	visited[i][j] = true
	if grid[i][j] == '0' {
		return
	}
	if i > 0 && grid[i-1][j] == '1' {
		dfs(grid, visited, i-1, j)
	}
	if i < len(grid)-1 && grid[i+1][j] == '1' {
		dfs(grid, visited, i+1, j)
	}
	if j > 0 && grid[i][j-1] == '1' {
		dfs(grid, visited, i, j-1)
	}
	if j < len(grid[0])-1 && grid[i][j+1] == '1' {
		dfs(grid, visited, i, j+1)
	}
}

func main() {
	grids := [][][]byte{
		{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		{
			{'1', '1', '1'},
			{'0', '1', '0'},
			{'1', '1', '1'},
		},
		{
			{'1', '0', '1', '1', '1'},
			{'1', '0', '1', '0', '1'},
			{'1', '1', '1', '0', '1'},
		},
		{
			{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '1', '1'},
			{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0'},
			{'1', '0', '1', '1', '1', '0', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1'},
			{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1'},
			{'1', '0', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '0'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		},
	}

	ans := []int{1, 3, 1, 1, 1}
	for i := 0; i < len(grids) && i < len(ans); i++ {
		fmt.Println("ret2:", numIslands2(grids[i]), "ans:", ans[i])
		fmt.Println("ret3:", numIslands3(grids[i]), "ans:", ans[i])
		fmt.Println("ret1:", numIslands(grids[i]), "ans:", ans[i]) // 会改变 grids
	}
}
