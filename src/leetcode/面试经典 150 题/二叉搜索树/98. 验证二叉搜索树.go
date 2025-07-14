package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func isValidBST(root *TreeNode) bool {
	arr = make([]int, 0, 16)
	inOrder(root)

	for i := len(arr) - 1; i >= 1; i-- {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	arr = append(arr, node.Val)
	if len(arr) > 1 {
		if arr[len(arr)-1] <= arr[len(arr)-2] { // 提前退出
			return
		}
	}
	inOrder(node.Right)
}
