package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyNode := &ListNode{
		Next: head,
	}

	pre := dummyNode
	cur := head
	next := cur.Next
	newCur := next.Next

	for cur != nil && next != nil {
		pre.Next = next
		next.Next = cur
		cur.Next = newCur

		pre = cur
		cur = newCur

		if cur == nil || cur.Next == nil {
			break
		}
		next = cur.Next
		newCur = next.Next
	}

	return dummyNode.Next
}
