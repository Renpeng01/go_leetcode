package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) >= 2 {
		l1 := merge(lists[0], lists[1])
		lists = lists[2:]
		lists = append(lists, l1)
	}
	return lists[0]
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	mhead := &ListNode{}
	dummy := mhead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			mhead.Next = l1
			l1 = l1.Next
		} else {
			mhead.Next = l2
			l2 = l2.Next
		}
		mhead = mhead.Next
	}

	if l1 != nil {
		mhead.Next = l1
	}
	if l2 != nil {
		mhead.Next = l2
	}
	return dummy.Next
}

func main() {
	// n1 := &ListNode{Val: 1}
	// n2 := &ListNode{Val: 2}
	// res := merge(n1, n2)

	// for res != nil {
	// 	fmt.Println(res.Val)
	// 	res = res.Next
	// }

	l := []int{1, 2, 3}

	for len(l) >= 2 {
		sum := l[0] + l[1]
		l = l[2:]
		l = append(l, sum)
	}

	fmt.Println(l)

}
