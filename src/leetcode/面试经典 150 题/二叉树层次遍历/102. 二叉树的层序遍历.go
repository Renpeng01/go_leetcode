package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 8)
	stack := make([]*TreeNode, 0, 16)
	stack = append(stack, root)
	for len(stack) > 0 {

		n := len(stack)
		levelVals := make([]int, 0, 8)
		for i := 0; i < n; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
			levelVals = append(levelVals, stack[i].Val)
		}
		stack = stack[n:]

		if len(levelVals) > 0 {
			res = append(res, levelVals)
		}

	}
	return res

}
