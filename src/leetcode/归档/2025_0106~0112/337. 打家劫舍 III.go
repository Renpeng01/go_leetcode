package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])

}

func robTree(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	leftDp := robTree(root.Left)
	rightDp := robTree(root.Right)

	robCur := root.Val + leftDp[0] + rightDp[0]
	notRobCur := max(leftDp[0], leftDp[1]) + max(rightDp[0], rightDp[1])

	return []int{notRobCur, robCur}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
