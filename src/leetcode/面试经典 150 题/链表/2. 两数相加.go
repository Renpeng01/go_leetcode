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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	t := 0
	resNode := &ListNode{}
	head := resNode
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + t
		t = sum / 10
		resNode.Next = &ListNode{
			Val: sum % 10,
		}
		resNode = resNode.Next
		l1 = l1.Next
		l2 = l2.Next

	}

	for l1 != nil {
		sum := l1.Val + t
		t = sum / 10
		resNode.Next = &ListNode{
			Val: sum % 10,
		}
		resNode = resNode.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + t
		t = sum / 10
		resNode.Next = &ListNode{
			Val: sum % 10,
		}
		resNode = resNode.Next
		l2 = l2.Next
	}

	if t != 0 {
		resNode.Next = &ListNode{
			Val: t,
		}
		resNode = resNode.Next

	}
	return head.Next
}
