package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head.Next
	head.Next = nil

	for cur != nil {
		temp := cur
		cur = cur.Next
		temp.Next = head
		head = temp
	}
	return head

}

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode = head
	var pre *ListNode = nil

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}

	return pre
}
