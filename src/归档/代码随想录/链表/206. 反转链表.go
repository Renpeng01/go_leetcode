package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	next := head.Next
	cur.Next = nil

	for next != nil {
		temp := next
		next = next.Next
		temp.Next = cur
		cur = temp
	}
	return cur
}
