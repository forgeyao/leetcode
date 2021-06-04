package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if right > depth {
		depth = right
	}
	return depth + 1
}

func main() {
	fmt.Println("")
}
