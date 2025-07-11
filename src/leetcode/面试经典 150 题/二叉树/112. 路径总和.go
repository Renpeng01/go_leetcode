package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hashChildPathSum(root, targetSum)
}

func hashChildPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	return hashChildPathSum(root.Left, targetSum-root.Val) || hashChildPathSum(root.Right, targetSum-root.Val)
}
