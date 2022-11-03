// https://leetcode.cn/problems/path-sum/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 时间 O(n) n为节点个数
 * 空间 O(h) h为树高度，最差情况为O(n)
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return (hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val))
}

func main() {
	n := [][]int{{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}, {1, 2, 3}, {1, 2}, {}}
	target := []int{22, 3, 0, 0}
	for i := 0; i < len(n) && i < len(target); i++ {
		fmt.Println("ret:", hasPathSum(util.CreateTree(n[i]), target[i]))
	}
}
