package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	res := pre
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		pre = pre.Next
		cur = cur.Next

	}
	return res.Next
}
