// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func inorderTraversal_recusion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	n = append(n, inorderTraversal_recusion(root.Left)...)
	n = append(n, root.Val)
	n = append(n, inorderTraversal_recusion(root.Right)...)
	return n
}

// 非递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	stack := []*TreeNode{}
	for node := root; node != nil; {
		if node.Left != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}
		n = append(n, node.Val)
		if node.Right != nil {
			node = node.Right
			continue
		}

		flag := false
		for len(stack) > 0 {
			node = stack[len(stack)-1]
			n = append(n, node.Val)
			stack = stack[:len(stack)-1]

			if node.Right != nil {
				node = node.Right
				flag = true
				break
			}
		}
		if !flag {
			break
		}
	}
	return n
}

// 非递归 更简洁的官方代码
func inorderTraversal_official(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	stack := []*TreeNode{}
	for node := root; node != nil || len(stack) > 0; {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}
		node = stack[len(stack)-1]
		n = append(n, node.Val)
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return n
}

func main() {
	n := [][]int{{1, 0, 2, 3}, {}, {1}, {1, 2}, {1, 0, 2}, {3, 1, 2}}
	for i := 0; i < len(n); i++ {
		var root, node *TreeNode
		array := []*TreeNode{}
		idx := 0
		for j := 0; j < len(n[i]); j++ {
			if n[i][j] == 0 {
				if j%2 == 0 {
					node = array[idx]
					idx++
				}
				continue
			}

			tmp := new(TreeNode)
			tmp.Val = n[i][j]
			if j == 0 {
				root = tmp
				node = root
			} else {
				if j%2 == 1 {
					node.Left = tmp
					array = append(array, tmp)
				} else {
					node.Right = tmp
					array = append(array, tmp)
				}
				if j%2 == 0 {
					node = array[idx]
					idx++
				}
			}
		}
		fmt.Println("recusion             :", inorderTraversal_recusion(root))
		fmt.Println("non-recusion         :", inorderTraversal(root))
		fmt.Println("non-recusion official:", inorderTraversal_official(root))
	}
}
