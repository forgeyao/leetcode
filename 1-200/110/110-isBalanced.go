//https://leetcode.cn/problems/balanced-binary-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 自顶向下递归
 * 时间 O(N*N)
 * 空间 O(N)
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := height(root.Left)
	r := height(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var getTreeLevel func(left *TreeNode, right *TreeNode) (int, bool)
	getTreeLevel = func(left *TreeNode, right *TreeNode) (int, bool) {
		if left == nil && right == nil {
			return 0, true
		}
		l, r := 0, 0
		b := true // 标记是否平衡
		if left != nil {
			l, b = getTreeLevel(left.Left, left.Right)
			l += 1
		}
		if right != nil && b {
			r, b = getTreeLevel(right.Left, right.Right)
			r += 1
		}
		if !b || l-r > 1 || l-r < -1 {
			b = false
		}

		if l < r {
			return r, b
		}
		return l, b
	}

	_, b := getTreeLevel(root.Left, root.Right)
	return b
}

/**
 * 自底向上递归
 * 时间 O(N)
 * 空间 O(N)
 */
func isBalanced3(root *TreeNode) bool {
	return height3(root) >= 0
}
func height3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := height3(root.Left)
	r := height3(root.Right)
	if l == -1 || r == -1 || abs(l-r) > 1 {
		return -1
	}
	return max(l, r) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 0, 3},
		{1, 2, 3, 4, 5, 0, 7, 8, 0, 0, 0, 0, 9},
		{3, 9, 20, 0, 0, 15, 7},
		{1, 2, 2, 3, 3, 0, 0, 4, 4},
	}
	for i := 0; i < len(n); i++ {
		root := util.CreateTree(n[i])
		fmt.Println("ret:", isBalanced(root), isBalanced3(root))
	}
}
