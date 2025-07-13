package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0, 128)
	queue = append(queue, root)

	stack := make([]*TreeNode, 0, 128)
	for len(queue) > 0 {
		node := queue[0]
		stack = append(stack, node)
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	cacheMax := make(map[*TreeNode]int, len(stack))
	maxPathSum := math.MinInt
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		left := 0
		if node.Left != nil {
			if v, ok := cacheMax[node.Left]; ok {
				left = v
			} else {
				left = maxPathSumFromLowLevel(node.Left)
			}
		}

		right := 0
		if node.Right != nil {
			if v, ok := cacheMax[node.Right]; ok {
				right = v
			} else {
				right = maxPathSumFromLowLevel(node.Right)
			}
		}

		maxPathSum = max(max(max(max(maxPathSum, node.Val), node.Val+right), node.Val+left), node.Val+left+right)
		cacheMax[node] = max(max(node.Val, node.Val+left), node.Val+node.Val+right)
	}

	return maxPathSum

}

// 包含node的最大路径和
func maxPathSumFromLowLevel(node *TreeNode) int {
	if node == nil {
		return 0
	}

	a := node.Val
	b := node.Val + maxPathSumFromLowLevel(node.Left)
	c := node.Val + maxPathSumFromLowLevel(node.Right)
	maxPathSum := max(max(a, b), c)
	return maxPathSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
