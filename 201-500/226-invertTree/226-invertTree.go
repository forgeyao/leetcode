// https://leetcode.cn/problems/invert-binary-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 递归(深度优先)
 * 时间 O(N)
 * 空间 平均 O(logN), 最差O(N)
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

/**
 * 非递归(广度优先)
 * 时间 O(N)
 * 空间 O(N)
 */
func invertTree2(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	for i := 0; i < len(queue); i++ {
		if queue[i] == nil {
			continue
		}
		queue[i].Left, queue[i].Right = queue[i].Right, queue[i].Left
		queue = append(queue, queue[i].Left, queue[i].Right)
	}
	return root
}

func main() {
	/*numss := [][]int{
		{},
		{2, 1, 3},
		{4, 2, 7, 1, 3, 6, 9},
	}*/
	fmt.Println("vim-go")
}
