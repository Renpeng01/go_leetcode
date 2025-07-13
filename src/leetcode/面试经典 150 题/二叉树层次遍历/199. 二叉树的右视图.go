package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 8)

	stack := make([]*TreeNode, 0, 16)
	stack = append(stack, root)

	for len(stack) > 0 {
		targetNode := stack[len(stack)-1]
		res = append(res, targetNode.Val)

		newStack := make([]*TreeNode, 0, 8)

		for _, v := range stack {
			if v.Left != nil {
				newStack = append(newStack, v.Left)
			}
			if v.Right != nil {
				newStack = append(newStack, v.Right)
			}
		}
		stack = newStack
	}

	return res
}
