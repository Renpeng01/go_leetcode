package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	curList1 := list1
	curList2 := list2

	cur := &ListNode{}
	head := cur

	for curList1 != nil && curList2 != nil {
		if curList1.Val < curList2.Val {
			cur.Next = curList1
			curList1 = curList1.Next

		} else {
			cur.Next = curList2
			curList1 = curList2.Next
		}
		cur = cur.Next
	}

	if curList1 != nil {
		cur.Next = curList1
	}

	if curList2 != nil {
		cur.Next = curList2
	}

	return head.Next
}
