package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	mHead := head
	cnt := 0

	for mHead != nil {
		cnt++
		mHead = mHead.Next
	}

	k = k % cnt
	if k == 0 {
		return head
	}

	fast := head
	slow := head

	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// fmt.Printf("fast: %+v slow: %+v\n", fast.Val, slow.Val)

	fast.Next = head
	res := slow.Next
	slow.Next = nil
	return res
}

func main() {
	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	res := rotateRight(node1, 2)
	i := 10
	for res != nil && i > 0 {
		fmt.Println("res:", res.Val)
		res = res.Next
		i--
	}

}
