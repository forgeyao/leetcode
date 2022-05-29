//https://leetcode-cn.com/problems/balanced-binary-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func isBalanced(root *TreeNode) bool {
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

func main() {
	n := [][]int{{3, 9, 20, 0, 0, 15, 7}, {1, 2, 2, 3, 3, 0, 0, 4, 4}}
	for i := 0; i < len(n); i++ {
		root := util.CreateTree(n[i])
		fmt.Println("ret:", isBalanced(root))
	}
}
