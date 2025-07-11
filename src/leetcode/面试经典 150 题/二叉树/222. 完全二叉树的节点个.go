package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0, 8)
	queue = append(queue, root)
	cnt := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil {
			cnt++
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			cnt++++
			queue = append(queue, node.Left)
		}
	}

	return cnt
}
