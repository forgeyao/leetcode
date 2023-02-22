// https://leetcode.cn/problems/binary-tree-preorder-traversal/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 递归
 * 时间O(N)
 * 空间 O(N), 平均O(logN), 最坏情况O(N)
 */
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

/**
 * 迭代
 * 时间 O(N)
 * 空间 O(N), 平均O(logN), 最坏情况O(N)
 */
func preorderTraversal2(root *TreeNode) []int {
	ret := []int{}
	stack := []*TreeNode{}
	for cur := root; cur != nil || len(stack) > 0; {
		for cur != nil {
			ret = append(ret, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		cur = cur.Right
		stack = stack[:len(stack)-1]
	}
	return ret
}

func main() {
	n := [][]int{
		{},
		{1},
		{1, 2},
		{3, 1, 2},
		{1, 0, 2},
		{1, 0, 2, 3},
		{2, 3, 0, 1},
		{3, 1, 0, 0, 2},
	}
	for i := 0; i < len(n); i++ {
		root := util.CreateTree(n[i])
		fmt.Println("ret:", preorderTraversal(root), preorderTraversal2(root))
	}
}
