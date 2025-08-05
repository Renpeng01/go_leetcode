package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将list转为array 超过了内存限制
func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l := make([]*ListNode, 0, 256)
	for head != nil {
		l = append(l, head)
		head = head.Next
	}

	res := mergeSort(l)

	mhead := res[0]
	result := mhead
	for i := 1; i < len(res); i++ {
		mhead.Next = res[i]
		mhead = mhead.Next
	}
	return result
}

func mergeSort(l []*ListNode) []*ListNode {
	if len(l) == 0 {
		return nil
	}
	mid := len(l) / 2
	l1 := mergeSort(l[0:mid])
	l2 := mergeSort(l[mid:])

	return merge(l1, l2)
}

func merge(l1, l2 []*ListNode) []*ListNode {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	res := make([]*ListNode, 0, len(l1)+len(l2))
	for len(l1) > 0 && len(l2) > 0 {
		if l1[0].Val < l2[0].Val {
			res = append(res, l1[0])
			l1 = l1[1:]
		} else {
			res = append(res, l2[0])
			l2 = l2[1:]
		}
	}

	if len(l1) > 0 {
		res = append(res, l1...)
	}
	if len(l2) > 0 {
		res = append(res, l2...)
	}
	return res
}
