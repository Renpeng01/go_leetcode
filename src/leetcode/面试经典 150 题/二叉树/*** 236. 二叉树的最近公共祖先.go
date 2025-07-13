package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	} else {
		return nil
	}
}

func main() {
	node0 := &TreeNode{
		Val: 3,
	}
	node1 := &TreeNode{
		Val: 5,
	}
	node2 := &TreeNode{
		Val: 1,
	}
	node0.Left = node1
	node0.Right = node2

	node3 := &TreeNode{
		Val: 6,
	}
	node4 := &TreeNode{
		Val: 2,
	}

	node1.Left = node3
	node1.Right = node4

	node5 := &TreeNode{
		Val: 0,
	}
	node6 := &TreeNode{
		Val: 8,
	}

	node2.Left = node5
	node2.Right = node6

	node7 := &TreeNode{
		Val: 7,
	}
	node8 := &TreeNode{
		Val: 4,
	}

	node4.Left = node7
	node4.Right = node8

	path1 := getNodePath(node0, node8)
	for _, v := range path1 {
		fmt.Println(v.Val)
	}
}
