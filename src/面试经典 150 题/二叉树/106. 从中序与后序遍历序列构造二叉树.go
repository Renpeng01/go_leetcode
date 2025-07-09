package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: postorder[len(postorder)-1],
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

	fmt.Println("leftInOrder:", leftInOrder, "rightInOrder:", rightInOrder)

	leftPostOrder := postorder[0:len(leftInOrder)]
	rightPostOrder := postorder[len(leftInOrder) : len(postorder)-1]

	fmt.Println("leftPostOrder:", leftPostOrder, "rightPostOrder:", rightPostOrder)

	root.Left = buildTree(leftInOrder, leftPostOrder)
	root.Right = buildTree(rightInOrder, rightPostOrder)
	return root

}
