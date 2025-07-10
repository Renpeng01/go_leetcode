package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head
	slow := head

	for fast.Next != nil {
		slow = slow.Next

		if fast.Next.Next == nil {
			break

		}
		fast = fast.Next.Next
	}
	return slow
}

func main() {

	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5

	res := middleNode(l1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
