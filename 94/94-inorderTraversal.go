// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	n = append(n, inorderTraversal(root.Left)...)
	n = append(n, root.Val)
	n = append(n, inorderTraversal(root.Right)...)
	return n
}

func main() {
	fmt.Println("vim-go")
}
