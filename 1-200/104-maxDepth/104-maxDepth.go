// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 深度优先搜索(递归)
 * 时间 O(N)
 * 空间 O(height) height 为二叉树高度
 */
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

/**
 * 广度优先搜索(迭代)
 * 时间 O(N)
 * 空间 O(N)
 */
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

func main() {
	nums := [][]int{
		{},
		{1},
		{1, 2},
		{1, 0, 2},
		{1, 2, 0, 3},
		{3, 9, 20, 0, 0, 15, 7},
		{2, 0, 3, 0, 4, 0, 5, 0, 6},
	}
	for i := 0; i < len(nums); i++ {
		root := util.CreateTree(nums[i])
		fmt.Println(nums[i], "ret:", maxDepth(root), maxDepth2(root))
	}
}
