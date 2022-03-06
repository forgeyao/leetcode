// https://leetcode-cn.com/problems/create-binary-tree-from-descriptions/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	m := map[int]*TreeNode{}
	mParent := map[int]*TreeNode{}

	for _, d := range descriptions {
		if _, ok := m[d[1]]; !ok {
			t := new(TreeNode)
			t.Val = d[1]
			t.Left = nil
			t.Right = nil
			m[d[1]] = t
		}

		node, ok := m[d[0]]
		if !ok {
			t := new(TreeNode)
			t.Val = d[0]
			if d[2] == 1 {
				t.Left = m[d[1]]
				t.Right = nil
			} else {
				t.Right = m[d[1]]
				t.Left = nil
			}
			m[d[0]] = t
			mParent[d[1]] = t
			node = t
		} else {
			if d[2] == 1 {
				node.Left = m[d[1]]
			} else {
				node.Right = m[d[1]]
			}
			mParent[d[1]] = node
		}
	}

	for _, v := range mParent {
		for v != nil {
			if mParent[v.Val] == nil {
				return v
			} else {
				v = mParent[v.Val]
			}
		}
		break
	}
	return nil
}
