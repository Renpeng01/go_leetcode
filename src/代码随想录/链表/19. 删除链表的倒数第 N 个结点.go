package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}

	fast := dummyNode
	for i := 0; i < n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}

	slow := dummyNode
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummyNode.Next

}
