package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodePtrMap := make(map[*Node]*Node, 256)

	loopHead := head
	copyHead := &Node{}
	res := copyHead

	for loopHead != nil {
		node := &Node{
			Val: loopHead.Val,
		}
		nodePtrMap[loopHead] = node
		copyHead.Next = node
		copyHead = copyHead.Next
		loopHead = loopHead.Next
	}

	setRandomNodeHead := res.Next
	for head != nil {
		if head.Random != nil {
			setRandomNodeHead.Random = nodePtrMap[head.Random]
		}
		head = head.Next
		setRandomNodeHead = setRandomNodeHead.Next
	}
	return res.Next
}
