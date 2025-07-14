package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0, 8)
	stack = append(stack, root)

	minDiff := math.MaxInt
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		a1 := math.MaxInt
		a2 := math.MaxInt

		if node.Left != nil {
			stack = append(stack, node.Left)
			a1 = node.Val - node.Left.Val
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			a2 = node.Right.Val - node.Val
		}
		minDiff = min(min(a1, a2), minDiff)
	}
	return minDiff
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

// 差值是一个正数，其数值等于两值之差的绝对值。
