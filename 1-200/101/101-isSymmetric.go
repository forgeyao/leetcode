// https://leetcode.cn/problems/symmetric-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 不能用中序遍历，再对比结果。因为结果相同，但区分不了顺序
 * 递归
 * 时间 O(n), n是节点个数
 * 空间 O(n)
 */
func isSymmetric(root *TreeNode) bool {
	var isSymmetric_cur func(left *TreeNode, right *TreeNode) bool
	isSymmetric_cur = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left != nil && right != nil && left.Val == right.Val {
			return isSymmetric_cur(left.Left, right.Right) && isSymmetric_cur(left.Right, right.Left)
		}
		return false
	}

	return isSymmetric_cur(root.Left, root.Right)
}

// 非递归
func isSymmetric_no(root *TreeNode) bool {
	if root == nil {
		return true
	}

	array := []*TreeNode{root.Left, root.Right}
	for len(array) > 0 {
		left := array[0]
		right := array[1]
		array = array[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		array = append(array, left.Left, right.Right)
		array = append(array, left.Right, right.Left)
	}
	return true
}

func main() {
	n := [][]int{{1, 2, 2, 3, 4, 4, 3}, {1, 2, 2, 0, 3, 0, 3}, {1, 2, 2, 2, 0, 2}, {1, 2, 2, 0, 3, 3}, {2, 3, 3, 4, 5, 5, 4, 0, 0, 8, 9, 0, 0, 9, 8}, {3, 4, 4, 5, 0, 0, 5, 6, 0, 0, 6}, {9, -42, -42, 0, 76, 76, 0, 0, 13, 0, 13}, {6, 82, 82, 0, 53, 53, 0, -58, 0, 0, -58, 0, -85, -85, 0, -9, 0, 0, -9, 0, 48, 48, 0, 33, 0, 0, 33, 81, 0, 0, 81, 5, 0, 0, 5, 61, 0, 0, 61, 0, 9, 9, 0, 91, 0, 0, 91, 72, 7, 7, 72, 89, 0, 94, 0, 0, 94, 0, 89, -27, 0, -30, 36, 36, -30, 0, -27, 50, 36, 0, -80, 34, 0, 0, 34, -80, 0, 36, 50, 18, 0, 0, 91, 77, 0, 0, 95, 95, 0, 0, 77, 91, 0, 0, 18, -19, 65, 0, 94, 0, -53, 0, -29, -29, 0, -53, 0, 94, 0, 65, -19, -62, -15, -35, 0, 0, -19, 43, 0, -21, 0, 0, -21, 0, 43, -19, 0, 0, -35, -15, -62, 86, 0, 0, -70, 0, 19, 0, 55, -79, 0, 0, -96, -96, 0, 0, -79, 55, 0, 19, 0, -70, 0, 0, 86, 49, 0, 25, 0, -19, 0, 0, 8, 30, 0, 82, -47, -47, 82, 0, 30, 8, 0, 0, -19, 0, 25, 0, 49}}
	for i := 0; i < len(n); i++ {
		root := util.CreateTree(n[i])
		//fmt.Println("tree:", util.PreorderTraversal(root))
		fmt.Println("ret1:", isSymmetric(root))
		fmt.Println("ret2:", isSymmetric_no(root))
	}
}
