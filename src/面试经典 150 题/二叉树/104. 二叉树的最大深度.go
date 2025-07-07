package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return maxChildDepth(root, 0)
}

func maxChildDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	return 1 + max(maxChildDepth(node.Left, depth), maxChildDepth(node.Right, depth))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
