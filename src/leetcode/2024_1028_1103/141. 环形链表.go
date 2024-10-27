package main

func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	path := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := path[head]; ok {
			return true
		}
		path[head] = struct{}{}
		head = head.Next
	}
	return false
}
