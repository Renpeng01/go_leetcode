package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Right)
	invertTree(root.Left)

	return root
}
