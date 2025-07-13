package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	stack := make([]*TreeNode, 0, 16)
	stack = append(stack, root)

	res := make([]float64, 0, 16)

	for len(stack) > 0 {
		sum := 0.0
		n := len(stack)
		for i := 0; i < n; i++ {
			sum += float64(stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		res = append(res, sum/float64(n))
		stack = stack[n:]
	}
	return res
}
