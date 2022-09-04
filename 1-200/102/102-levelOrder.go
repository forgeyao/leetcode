// https://leetcode.cn/problems/binary-tree-level-order-traversal/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

/**
 * 从根节点开始入栈，左右存在的子节点入栈
 * 栈顶出栈，依次迭代只到栈为空
 * 即广度优先搜索
 *
 * 时间 O(n), n为节点个数
 * 空间 O(n)
 */
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{root} // 节点栈
	levSize := []int{1}        // 每层节点数
	for i, idxStack := 0, 0; i < len(levSize); i++ {
		r := []int{} // 当前层输出
		num := 0     // 下一层节点数
		for j := 0; j < levSize[i]; j++ {
			idx := stack[j+idxStack]
			r = append(r, idx.Val)
			if idx.Left != nil {
				stack = append(stack, idx.Left)
				num++
			}
			if idx.Right != nil {
				stack = append(stack, idx.Right)
				num++
			}
		}
		if len(r) > 0 {
			levSize = append(levSize, num)
			ret = append(ret, r)
		}
		idxStack += levSize[i]
		//fmt.Println(r, num, len(stack))
	}
	return ret
}

/**
 * 官方答案。更简洁
 */
func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{root}
	for i := 0; 0 < len(stack); i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(stack); j++ {
			idx := stack[j]
			ret[i] = append(ret[i], idx.Val)
			if idx.Left != nil {
				p = append(p, idx.Left)
			}
			if idx.Right != nil {
				p = append(p, idx.Right)
			}
		}
		stack = p
	}
	return ret
}

func main() {
	roots := [][]int{
		{3, 9, 20, 0, 0, 15, 7},
		{1},
		{},
	}
	ans := [][][]int{
		{{3}, {9, 20}, {15, 7}},
		{{1}},
		{},
	}
	for i := 0; i < len(roots); i++ {
		//fmt.Println("ret:", levelOrder(util.CreateTree(roots[i])), "ans:", ans[i])
		fmt.Println("ret:", levelOrder2(util.CreateTree(roots[i])), "ans:", ans[i])
	}
}
