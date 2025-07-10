package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}
	stack := make([]*Node, 0, 256)

	if root.Left != nil {
		stack = append(stack, root.Left)
	}

	if root.Right != nil {
		stack = append(stack, root.Right)
	}
	for len(stack) > 0 {
		newStack := make([]*Node, 0, 256)
		if stack[0].Left != nil {
			newStack = append(newStack, stack[0].Left)
		}
		if stack[0].Right != nil {
			newStack = append(newStack, stack[0].Right)
		}
		for i := 1; i < len(stack); i++ {
			stack[i-1].Next = stack[i]

			if stack[i].Left != nil {
				newStack = append(newStack, stack[i].Left)
			}
			if stack[i].Right != nil {
				newStack = append(newStack, stack[i].Right)
			}
		}
		stack = newStack
	}
	return root
}
