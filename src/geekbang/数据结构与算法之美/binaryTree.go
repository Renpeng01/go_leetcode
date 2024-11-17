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

// 层序遍历TBD

// 二叉查找树： 在树中的任意节点，其左子树中的每个节点的值，都小于这个节点的值，而右子树节点的值都大于这个节点的值
// 支持动态数据集合的快速插入，删除和查找

type BinarySearchTree struct {
	Root *BinaryTreeNode
}

func (binarySearchTree *BinarySearchTree) Find(data int) *BinaryTreeNode {
	p := binarySearchTree.Root
	for p != nil {
		if data < p.Val {
			p = p.Left
		} else if data > p.Val {
			p = p.Right
		} else {
			return p
		}
	}
	return nil
}

func (binarySearchTree *BinarySearchTree) Insert(data int) {
	if binarySearchTree.Root == nil {
		binarySearchTree.Root = &BinaryTreeNode{
			Val: data,
		}
		return
	}

	p := binarySearchTree.Root
	for p != nil {
		if data > p.Val {
			if p.Right == nil {
				p.Right = &BinaryTreeNode{Val: data}
				return
			}
			p = p.Right
		} else {
			if p.Left == nil {
				p.Left = &BinaryTreeNode{Val: data}
				return
			}
			p = p.Left
		}
	}
}

func (binarySearchTree *BinarySearchTree) Delete(data int) {
	p := binarySearchTree.Root
	var pp *BinaryTreeNode
	for p != nil && p.Val != data {
		pp = p
		if data > p.Val {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	if p == nil {
		return
	}

	// 要删除的节点右2个子节点
	if p.Left != nil && p.Right != nil { // 查找右子树中最小节点
		minP := p.Right
		minPP := p
		for minP.Left != nil {
			minPP = minP
			minP = minP.Left
		}
		p.Val = minP.Val // 将minP的数据替换到p中
		p = minP         // 下面就变成了删除minP了
		pp = minPP
	}

	var child *BinaryTreeNode
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	}

	if pp == nil {
		binarySearchTree.Root = child
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
}
