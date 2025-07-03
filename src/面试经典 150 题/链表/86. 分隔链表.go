package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	pre := &ListNode{}
	next := &ListNode{}

	mPre := pre
	mNext := next

	for head != nil {
		if head.Val < x {
			mPre.Next = head
			mPre = mPre.Next

		} else if head.Val > x {
			mNext.Next = head
			mNext = mNext.Next
		} else {
			if next.Next != nil || pre.Next == nil {
				mNext.Next = head
				mNext = mNext.Next
			} else {
				mPre.Next = head
				mPre = mPre.Next
			}
		}

		h := head
		head = head.Next
		h.Next = nil

	}

	// log
	logPre := pre
	logNext := next
	for logPre.Next != nil {
		fmt.Println("pre:", logPre.Next.Val)
		logPre = logPre.Next
	}

	for logNext.Next != nil {
		fmt.Println("next:", logNext.Next.Val)
		logNext = logNext.Next
	}

	if pre.Next == nil {
		return next.Next
	}

	if next.Next == nil {
		return pre.Next
	}

	mPre.Next = next.Next
	return pre.Next
}

func main() {
	// node1 := &ListNode{Val: 1, Next: nil}
	// node2 := &ListNode{Val: 4, Next: nil}
	// node3 := &ListNode{Val: 3, Next: nil}
	// node4 := &ListNode{Val: 2, Next: nil}
	// node5 := &ListNode{Val: 5, Next: nil}
	// node6 := &ListNode{Val: 2, Next: nil}

	// node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4
	// node4.Next = node5
	// node5.Next = node6

	// node1 := &ListNode{Val: 2, Next: nil}
	// node2 := &ListNode{Val: 1, Next: nil}
	// node1.Next = node2

	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 3, Next: nil}
	node3 := &ListNode{Val: 2, Next: nil}
	node1.Next = node2
	node2.Next = node3

	// node1 := &ListNode{Val: 1, Next: nil}

	mHead := node1
	for mHead != nil {
		fmt.Println("origin: ", mHead.Val)
		mHead = mHead.Next
	}

	// partition(node1, 3)

	res := partition(node1, 3)
	i := 10
	for res != nil && i > 0 {
		fmt.Println("res:", res.Val)
		res = res.Next
		i--
	}

}
