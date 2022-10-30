// https://leetcode.cn/problems/diameter-of-binary-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 思路：最大直径有三种可能，左节点最大直径、右节点最大直径、左右节点和
 * 时间 O(n), n 为节点数
 * 空间 O(h), h 为树高度
 */
func diameterOfBinaryTree(root *TreeNode) int {
	l, m1 := deep(root.Left)
	r, m2 := deep(root.Right)
	return max(l+r, m1, m2)
}

func deep(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left, m1 := deep(root.Left)
	right, m2 := deep(root.Right)
	m := max(left+right, m1, m2)
	if left > right {
		return left + 1, m
	}
	return right + 1, m
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

/**
 * 参考官方答案，用类似全局变量方式，减少一个返回值
 */
func diameterOfBinaryTree2(root *TreeNode) int {
	max := 0
	deep2(root, &max)
	return max - 1
}

func deep2(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := deep2(root.Left, max)
	right := deep2(root.Right, max)
	if left+right+1 > *max {
		*max = left + right + 1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	nums := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 0, 3, 4, 0, 0, 5},
	}
	results := []int{3, 3}
	for i := 0; i < len(nums) && i < len(results); i++ {
		root := util.CreateTree(nums[i])
		fmt.Println("ret:", diameterOfBinaryTree(root), "result:", results[i])
		fmt.Println("ret:", diameterOfBinaryTree2(root), "result:", results[i])
	}
}
