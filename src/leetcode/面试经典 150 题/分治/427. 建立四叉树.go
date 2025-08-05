package main

import "fmt"

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

	isAll := false
	for i := top;i <= bottom;

	horizontalMid := left + (right-left)/2
	verticalMid := top + (bottom-top)/2

	root := &Node{}

	// fmt.Println("TopLeft", left, horizontalMid, top, verticalMid)
	root.TopLeft = build(grid, left, horizontalMid, top, verticalMid)

	// fmt.Println("TopRight", horizontalMid+1, right, top, verticalMid)
	root.TopRight = build(grid, horizontalMid+1, right, top, verticalMid)

	// fmt.Println("BottomLeft", left, horizontalMid, verticalMid+1, bottom)
	root.BottomLeft = build(grid, left, horizontalMid, verticalMid+1, bottom)

	// fmt.Println("BottomRight", horizontalMid+1, right, verticalMid+1, bottom)
	root.BottomRight = build(grid, horizontalMid+1, right, verticalMid+1, bottom)
	return root
}

func main() {
	grid := [][]int{{0, 1}, {1, 0}}
	res := construct(grid)

	fmt.Println(res.TopLeft)
	fmt.Println(res.TopRight)
	fmt.Println(res.BottomLeft)
	fmt.Println(res.BottomRight)
}
