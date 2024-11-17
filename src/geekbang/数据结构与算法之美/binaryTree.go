package main

import "fmt"

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// 遍历二叉树

// 前序遍历：先打印这个节点，然后再打印它的左子树，最后打印右子树
func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("val: %+v\n", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// 中序遍历：先打印它的左子树，然后再打印它本身，最后打印它的右子树
func inOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf("val: %+v\n", root.Val)
	inOrder(root.Right)
}

// 后序遍历：先打印左子树，在打印它的右子树，最后打印这个节点本身
func postOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Printf("val: %+v\n", root.Val)
}
