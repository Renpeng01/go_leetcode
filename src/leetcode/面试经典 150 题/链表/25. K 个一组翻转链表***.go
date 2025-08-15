package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	dummy := &ListNode{
		Next: head,
	}
	groupPre := dummy

	for {

		mk := k
		cur := groupPre
		for mk > 0 && cur != nil {
			cur = cur.Next
			mk--
		}
		if cur == nil {
			break
		}
		groupNext := cur.Next

		pre := groupNext
		cur = groupPre.Next

		for cur != groupNext {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}

		tmp := groupPre.Next
		groupPre.Next = pre
		groupPre = tmp
	}

	return dummy.Next
}
