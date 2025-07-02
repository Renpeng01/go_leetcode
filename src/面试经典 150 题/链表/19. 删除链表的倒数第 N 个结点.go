package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	sentinel := &ListNode{
		Next: head,
	}

	fast := sentinel
	slow := sentinel
	res := sentinel

	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}

	if n > 0 {
		return res.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return res.Next

}
