package main

import (
	"fmt"
	"strconv"
)

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

	res := mergeSort1(l)

	mhead := res[0]
	result := mhead
	for i := 1; i < len(res); i++ {
		mhead.Next = res[i]
		mhead = mhead.Next
	}
	return result
}

func mergeSort1(l []*ListNode) []*ListNode {
	if len(l) == 0 {
		return nil
	}
	mid := len(l) / 2
	l1 := mergeSort1(l[0:mid])
	l2 := mergeSort1(l[mid:])

	return merge1(l1, l2)
}

func merge1(l1, l2 []*ListNode) []*ListNode {
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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for slow != nil && slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	leftHead := head
	rightHead := slow.Next
	slow.Next = nil

	mLeftHead := leftHead

	mLeftStr := ""
	for mLeftHead != nil {
		mLeftStr += strconv.Itoa(mLeftHead.Val) + ","
		mLeftHead = mLeftHead.Next
	}
	fmt.Println("mLeftStr: ", mLeftStr)

	mRightHead := rightHead
	mRightStr := ""
	for mRightHead != nil {
		mRightStr += strconv.Itoa(mRightHead.Val) + ","
		mRightHead = mRightHead.Next
	}
	fmt.Println("mLeftStr: ", mRightStr)

	left := sortList(leftHead)
	right := sortList(rightHead)
	return merge(left, right)

}

func merge(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var mhead *ListNode
	if left.Val < right.Val {
		mhead = left
		left = left.Next
	} else {
		mhead = right
		right = right.Next
	}
	res := mhead
	for left != nil && right != nil {
		if left.Val < right.Val {
			mhead.Next = left
			left = left
		} else {
			mhead.Next = right
			right = right.Next
		}
		mhead = mhead.Next
	}

	if left != nil {
		mhead.Next = left
	}

	if right != nil {
		mhead.Next = right
	}
	return res
}

func main() {
	// [4,2,1,3]
	n1 := &ListNode{
		Val: 4,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 1,
	}
	n4 := &ListNode{
		Val: 3,
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	res := sortList(n1)

	resStr := ""
	for res != nil {
		resStr += strconv.Itoa(res.Val)
		res = res.Next
	}
	fmt.Println("--------------")
	fmt.Println(resStr)
}
