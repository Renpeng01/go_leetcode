package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	s     *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	if root == nil {
		return it
	}

	stack := make([]*TreeNode, 0, 8)
	cur := root
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}

	s := &TreeNode{}
	stack[len(stack)-1].Left = s
	it.s = s
	it.stack = stack
	return it
}

func (this *BSTIterator) Next() int {

	this.stack = this.stack[1:]
	node := this.stack[0]

	cur := node.Right
	if cur != nil {
		this.stack = append(this.stack, cur)
		cur = cur.Left
	}
	return node.Val

}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 1
}

func main() {
	node1 := &TreeNode{
		Val: 7,
	}
	node2 := &TreeNode{
		Val: 3,
	}
	node3 := &TreeNode{
		Val: 15,
	}
	node4 := &TreeNode{
		Val: 9,
	}
	node5 := &TreeNode{
		Val: 20,
	}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	it := Constructor(node1)

	a1 := it.Next()
	fmt.Println("a1: ", a1)

	a2 := it.Next()
	fmt.Println("a2: ", a2)

	a3 := it.HasNext()
	fmt.Println("a3: ", a3)

	a4 := it.Next()
	fmt.Println("a4: ", a4)

	a5 := it.HasNext()
	fmt.Println("a5: ", a5)

	a6 := it.Next()
	fmt.Println("a6: ", a6)

	a7 := it.HasNext()
	fmt.Println("a7: ", a7)

	a8 := it.Next()
	fmt.Println("a8: ", a8)

	a9 := it.HasNext()
	fmt.Println("a9: ", a9)

}
