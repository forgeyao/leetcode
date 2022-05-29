// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l == 0 {
		return r + 1
	}
	if r == 0 {
		return l + 1
	}
	if l > r {
		return r + 1
	}
	return l + 1
}

func main() {
	n := [][]int{{3, 9, 20, 0, 0, 15, 7}, {2, 0, 3, 0, 4, 0, 5, 0, 6}}
	for i := 0; i < len(n); i++ {
		root := util.CreateTree(n[i])
		fmt.Println("ret:", minDepth(root))
	}
}
