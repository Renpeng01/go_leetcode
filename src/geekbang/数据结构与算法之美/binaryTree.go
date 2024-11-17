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

// 支持重复数据的二叉查找树
// 1. 通过链表和支持动态扩容的数组等数据结构，把值相同的数据存储在通一个节点上
// 2. 每个节点仍只存储一个数据，在查找插入位置的过程中，如果碰到一个节点的值，与要插入数据的值相同，就将这个要插入的数据放到这个节点的右子树，把这个新插入的数据当作大于这个节点的值来处理。在查找数据时，遇到值相同的节点，我们并不停止查找操作，而是继续在右子树查找，直到遇到叶子节点。这样可以把键值等于要查找值的所有节点都找出来
// 对于删除操作，也需要找到每个要删除的节点一次删除

// 二叉查找树与散列表的优势
// 1. 散列表中的数据是无序存储的，如果要输出有序的数据，需要先进行排序，而对于二叉查找树，只需要中序遍历，可以在O(n)的时间复杂度内，输出有序的数据序列
// 2. 散列表扩容耗时，而且当遇到散列冲突时，性能不稳定
// 3. 尽管散列表的查找等操作的时间复杂度是常量级的，但因为哈希冲突的存在，这个常量级不一定比O(logn)小
// 4. 散列表的构造比二叉查找树复杂 如 哈希函数设计，冲突解决，扩容，缩容
