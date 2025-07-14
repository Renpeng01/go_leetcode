package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	arr = make([]int, 0, 16)

	if len(arr) == 1 {
		return arr[0]
	}
	inOrder(root)

	min := math.MaxInt
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
	}
	return min
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	arr = append(arr, node.Val)
	inOrder(node.Right)
}
