package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	setA := make(map[*ListNode]struct{}, 256)
	for headA != nil {
		setA[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := setA[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// https://www.bilibili.com/video/BV1xS4y1v7pp/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
