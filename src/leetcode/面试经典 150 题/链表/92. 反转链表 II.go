package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if left == right {
		return head
	}

	sentinel := &ListNode{
		Next: head,
	}

	beforeLeftNode := sentinel
	rightNode := sentinel
	mLeft := left
	mRight := right

	for mLeft > 1 {
		beforeLeftNode = beforeLeftNode.Next
		mLeft--
	}
	// fmt.Printf("beforeLeftNode: %+v\n", beforeLeftNode.Val)

	rightNode = beforeLeftNode.Next
	for mRight-left > 0 {
		rightNode = rightNode.Next
		mRight--
	}
	mRightNode := rightNode.Next
	// fmt.Printf("rightNode1: %+v\n", rightNode.Val)

	pre := beforeLeftNode.Next
	mPre := pre
	curNode := pre.Next
	next := curNode.Next
	pre.Next = nil

	for pre != rightNode {
		curNode.Next = pre
		pre = curNode
		curNode = next

		if next == nil {
			break
		}
		next = next.Next
	}

	beforeLeftNode.Next = rightNode
	mPre.Next = mRightNode
	// fmt.Printf("rightNode2: %+v\n", rightNode.Val)
	// fmt.Printf("mPre: %+v\n", mPre.Val)
	// fmt.Printf("rightNode.Next: %+v\n", rightNode.Next.Val)
	// fmt.Printf("mRightNode.Next: %+v\n", mRightNode.Next.Val)
	return sentinel.Next
}

func main() {
	// node1 := &ListNode{Val: 1, Next: nil}
	// node2 := &ListNode{Val: 2, Next: nil}
	// node3 := &ListNode{Val: 3, Next: nil}
	// node4 := &ListNode{Val: 4, Next: nil}
	// node5 := &ListNode{Val: 5, Next: nil}

	// node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4
	// node4.Next = node5
	// // reverseBetween(node1, 2, 4)
	// res := reverseBetween(node1, 2, 4)

	node1 := &ListNode{Val: 3, Next: nil}
	node2 := &ListNode{Val: 5, Next: nil}
	node1.Next = node2
	res := reverseBetween(node1, 1, 2)

	i := 10
	for res != nil && i > 0 {
		fmt.Println("res:", res.Val)
		res = res.Next
		i--
	}

}
