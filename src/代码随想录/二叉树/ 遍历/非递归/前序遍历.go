package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 16)
	stack := make([]*TreeNode, 0, 16)
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return res
}
