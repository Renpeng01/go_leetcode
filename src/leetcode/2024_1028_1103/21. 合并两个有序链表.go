package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	curList1 := list1
	curList2 := list2

	var newList *ListNode = &ListNode{}
	var head *ListNode = &ListNode{Next: newList}

	for curList1 != nil && curList2 != nil {
		if curList1.Val < curList2.Val {
			newList.Next = curList1
			curList1 = curList1.Next
		} else {
			newList.Next = curList2
			curList2 = curList2.Next
		}
		newList = newList.Next
	}
	for curList1 != nil {
		newList.Next = curList1
		newList = newList.Next
		curList1 = curList1.Next
	}

	for curList2 != nil {
		newList.Next = curList2
		newList = newList.Next
		curList2 = curList2.Next
	}

	return head.Next.Next

}
