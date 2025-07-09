package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {

		cur := root.Left
		for cur.Right != nil {
			cur = cur.Right
		}

		cur.Right = root.Right
		root.Right = root.Left
		root.Left = nil
		flatten(root.Right)
	}
}
