//https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}

	mid := l / 2
	root := new(TreeNode)
	root.Val = nums[mid]
	if mid >= 1 {
		root.Left = sortedArrayToBST(nums[0:mid])
	}
	if mid < l-1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}

func main() {
	n := [][]int{{1}, {-10, -3, 0, 5, 9}, {1, 3}, {1, 2, 3, 4, 5, 6, 7}}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", util.InorderTraversal(sortedArrayToBST(n[i])))
	}
}
