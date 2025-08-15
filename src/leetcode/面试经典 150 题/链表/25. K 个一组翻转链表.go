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
		cur := groupPre
		for k > 0 && cur != nil {
			cur = cur.Next
			k--
		}
		if cur == nil {
			break
		}
		groupNext := cur.Next

		pre := groupNext
		cur = groupPre.Next
		// var tmp *ListNode

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
