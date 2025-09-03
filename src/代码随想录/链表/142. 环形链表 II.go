package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 常考
// https://www.bilibili.com/video/BV1if4y1d7ob/?vd_source=70c464e99440c207e9933663bb2e5166
func detectCycle(head *ListNode) *ListNode {
	dummy := head
	fast := head
	slow := head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			for dummy != slow {
				dummy = dummy.Next
				slow = slow.Next
			}
			return dummy

		}
	}

	return nil
}
