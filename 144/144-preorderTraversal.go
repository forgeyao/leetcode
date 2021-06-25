// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func preorderTraversal(root *TreeNode) []int {
	num := []int{}
	if root == nil {
		return num
	}
	num = append(num, root.Val)
	num = append(num, preorderTraversal(root.Left)...)
	num = append(num, preorderTraversal(root.Right)...)
	return num
}

func main() {
	n := [][]int{{1, 0, 2, 3}, {}, {1}, {1, 2}, {1, 0, 2}}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", preorderTraversal(util.CreateTree(n[i])))
	}
}
