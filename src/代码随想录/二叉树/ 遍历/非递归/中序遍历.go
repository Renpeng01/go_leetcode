package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	res := make([]int, 0, 16)
	queue := make([]*TreeNode, 0, 16)
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, node.Left)
			continue
		}
		res = append(res, node.Val)
		if node.Right != nil {
			queue = append([]*TreeNode{node.Right}, queue...)
			continue
		}
		queue = queue[1:]
	}
	return res
}
