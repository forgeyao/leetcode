// https://leetcode-cn.com/problems/successor-lcci/
package main

import (
	"fmt"
	"leetcode/util"
)

type TreeNode = util.TreeNode

//  递归(深度优先搜索)
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	ret := inorder(root)
	for i := 0; i < len(ret); i++ {
		if ret[i].Val == p.Val {
			if i+1 < len(ret) {
				return ret[i+1]
			} else {
				return nil
			}
		}
	}
	return nil
}
func inorder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	ret := inorder(root.Left)
	ret = append(ret, root)
	ret = append(ret, inorder(root.Right)...)
	return ret
}

//  非递归(广度优先搜索)
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	head := root
	find := false
	for head != nil || len(stack) > 0 {
		for ; head != nil; head = head.Left {
			stack = append(stack, head)
		}
		top := stack[len(stack)-1]
		if find {
			return top
		}
		if top.Val == p.Val {
			find = true
		}
		head = top.Right
		stack = stack[:len(stack)-1]
	}
	return nil
}

func main() {
	n := [][]int{{2, 1, 3}, {5, 3, 6, 2, 4, 0, 0, 1}, {6, 2, 8, 0, 4, 7, 9, 0, 0, 3, 5}}
	//k := []int{1, 6, 2}
	for i := 0; i < len(n); /*&& i < len(k)*/ i++ {
		for j := 0; j < len(n[i]); j++ {
			p := new(TreeNode)
			//p.Val = k[i]
			p.Val = n[i][j]
			fmt.Println("ret:", n[i][j], "->", inorderSuccessor(util.CreateTree(n[i]), p))
			fmt.Println("ret:", n[i][j], "->", inorderSuccessor2(util.CreateTree(n[i]), p))
		}
	}
}
