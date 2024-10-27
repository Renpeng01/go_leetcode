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
func reverseList(head *ListNode) *ListNode {
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
