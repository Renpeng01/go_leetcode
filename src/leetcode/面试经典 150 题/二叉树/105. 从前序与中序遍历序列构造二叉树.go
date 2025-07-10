package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	rootIndex := 0
	for i, v := range inorder {
		if v == root.Val {
			rootIndex = i
			break
		}
	}
	leftInOrder := inorder[0:rootIndex]
	rightInOrder := inorder[rootIndex+1:]

	leftPreOrder := preorder[1 : 1+len(leftInOrder)]
	rightPreOrder := preorder[len(leftInOrder)+1:]

	root.Left = buildTree(leftPreOrder, leftInOrder)
	root.Right = buildTree(rightPreOrder, rightInOrder)
	return root
}
