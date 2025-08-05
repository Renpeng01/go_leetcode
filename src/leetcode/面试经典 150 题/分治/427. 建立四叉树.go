package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return build(grid, 0, len(grid)-1, 0, len(grid[0])-1)
}

func build(grid [][]int, left, right, top, bottom int) *Node {
	if left == right && top == bottom {
		node := &Node{}
		if grid[left][top] == 1 {
			node.Val = true
			node.IsLeaf = true
		}
		return node
	}

	horizontalMid := left + (right-left)/2
	verticalMid := top + (bottom-top)/2

	root := &Node{}
	root.TopLeft = build(grid, left, horizontalMid, top, verticalMid)
	root.TopRight = build(grid, horizontalMid+1, right, top, verticalMid)
	root.BottomLeft = build(grid, left, horizontalMid, verticalMid+1, bottom)
	root.BottomRight = build(grid, horizontalMid+1, right, verticalMid+1, bottom)
	return root
}
