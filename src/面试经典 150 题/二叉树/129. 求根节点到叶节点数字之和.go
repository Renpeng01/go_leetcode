package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := new(int)
	nums(root, 0, sum)
	return *sum
}

func nums(node *TreeNode, preNum int, sum *int) {
	if node == nil {
		return
	}
	preNum = preNum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*sum += preNum
		return
	}

	if node.Left != nil {
		nums(node.Left, preNum, sum)
	}

	if node.Right != nil {
		nums(node.Right, preNum, sum)
	}
}

func main() {

	root := &TreeNode{
		Val: 1,
	}

	left := &TreeNode{
		Val: 2,
	}
	right := &TreeNode{
		Val: 3,
	}
	root.Left = left
	root.Right = right

	res := sumNumbers(root)
	fmt.Println(res)
}
