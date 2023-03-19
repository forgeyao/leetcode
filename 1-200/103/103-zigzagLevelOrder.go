// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 分层迭代, 算是广度优先变种
 * 时间 O(N)
 * 空间 O(N)
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	right := false
	l := 1
	ret := [][]int{}
	for i := 0; i < len(queue); {
		tmp := []int{}
		tmpQueue := []*TreeNode{}
		for j := 0; j < l && i < len(queue); i, j = i+1, j+1 {
			tmp = append(tmp, queue[i].Val)
			if right {
				if queue[i].Right != nil {
					tmpQueue = append(tmpQueue, queue[i].Right)
				}
				if queue[i].Left != nil {
					tmpQueue = append(tmpQueue, queue[i].Left)
				}
			} else {
				if queue[i].Left != nil {
					tmpQueue = append(tmpQueue, queue[i].Left)
				}
				if queue[i].Right != nil {
					tmpQueue = append(tmpQueue, queue[i].Right)
				}
			}
		}
		for k := len(tmpQueue) - 1; k >= 0; k-- {
			queue = append(queue, tmpQueue[k])
		}
		right = !right
		l *= 2
		ret = append(ret, tmp)
	}
	return ret
}

func main() {
	numss := [][]int{
		{},
		{1},
		{3, 9, 20, 0, 0, 15, 7},
		{1, 2, 3, 4, 0, 0, 5},
	}
	for _, nums := range numss {
		root := util.CreateTree(nums)
		fmt.Println(nums, "ret:", zigzagLevelOrder(root))
	}
}
