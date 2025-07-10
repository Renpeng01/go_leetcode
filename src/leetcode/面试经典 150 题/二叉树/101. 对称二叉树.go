package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isSymmetricChild(root.Left, root.Right)

}

func isSymmetricChild(node1, node2 *TreeNode) bool {

	if node1 == nil && node2 == nil {
		return true
	}

	if node1 != nil && node2 != nil && node1.Val == node2.Val {
		return isSymmetricChild(node1.Left, node2.Right) && isSymmetricChild(node1.Right, node2.Left)
	}

	return false
}
