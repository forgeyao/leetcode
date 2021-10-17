package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func InorderTraversal(k int, root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	n = append(n, InorderTraversal(k, root.Left)...)
	n = append(n, root.Val)
	if len(n) >= k {
		return n
	}
	n = append(n, InorderTraversal(k, root.Right)...)
	return n
}

func kthSmallest(root *TreeNode, k int) int {
	n := InorderTraversal(k, root)
	return n[k-1]
}

func main() {
	n := [][]int{{3, 1, 4, 0, 2}, {1}, {5, 3, 6, 2, 4, 0, 0, 1}}
	k := []int{1, 1, 3}
	for i := 0; i < len(n) && i < len(k); i++ {
		fmt.Println("ret:", kthSmallest(util.CreateTree(n[i]), k[i]))
	}
}
