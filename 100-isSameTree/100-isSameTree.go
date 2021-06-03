// https://leetcode-cn.com/problems/same-tree/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	// 前序遍历
	var preOrder func(t1 *TreeNode, t2 *TreeNode) bool
	preOrder = func(t1 *TreeNode, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil || t1.Val != t2.Val {
			return false
		}

		return !preOrder(t1.Left, t2.Left) && !preOrder(t1.Right, t2.Right)
	}

	return preOrder(p, q)
}

func main() {
	fmt.Println("")
	/*trees := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2}, {1, nil, 2}, {1, 2, 1}, {1, 1, 2}}
	for i := 0; i < len(trees); i++ {
	}*/
}
