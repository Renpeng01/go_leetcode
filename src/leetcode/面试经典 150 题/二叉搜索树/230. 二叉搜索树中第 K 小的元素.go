package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	arr = make([]int, 0, 16)

	if len(arr) == 1 {
		return arr[0]
	}
	inOrder(root)

	return arr[k-1]

}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	arr = append(arr, node.Val)
	inOrder(node.Right)
}
