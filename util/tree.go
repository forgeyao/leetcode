package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(n []int) *TreeNode {
	var root, node *TreeNode
	array := []*TreeNode{}
	idx := 0
	for j := 0; j < len(n); j++ {
		if n[j] == 0 {
			if j%2 == 0 {
				node = array[idx]
				idx++
			}
			continue
		}

		tmp := new(TreeNode)
		tmp.Val = n[j]
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
	return root
}

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	n = append(n, root.Val)
	n = append(n, PreorderTraversal(root.Left)...)
	n = append(n, PreorderTraversal(root.Right)...)
	return n
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	n = append(n, InorderTraversal(root.Left)...)
	n = append(n, root.Val)
	n = append(n, InorderTraversal(root.Right)...)
	return n
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	n := []int{}
	n = append(n, PostorderTraversal(root.Left)...)
	n = append(n, PostorderTraversal(root.Right)...)
	n = append(n, root.Val)
	return n
}
