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
	inOrder(root, k)

	return arr[k-1]

}

func inOrder(node *TreeNode, k int) {
	if node == nil {
		return
	}
	inOrder(node.Left, k)
	arr = append(arr, node.Val)
	if len(arr) == k { // 这里是关键，需要提前结束，否则性能不好
		return
	}
	inOrder(node.Right, k)
}
