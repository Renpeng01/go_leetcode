package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sentinel := &ListNode{
		Next: head,
	}

	pre := sentinel
	cur := head
	next := cur.Next

	for {
		if next == nil {
			break
		}
		if next.Val != cur.Val {
			pre = cur
			cur = next
			next = next.Next
			continue
		}
		for next != nil {
			if next.Val != cur.Val {
				break
			}
			next = next.Next
		}
		pre.Next = next
		if next == nil || next.Next == nil {
			break
		}

		cur = next
		next = next.Next
	}

	return sentinel.Next
}
