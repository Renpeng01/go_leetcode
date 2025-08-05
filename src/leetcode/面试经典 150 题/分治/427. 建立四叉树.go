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
	if len(grid) == 1 {
		node := &Node{}
		if grid[0][0] == 1 {
			node.Val = true
			node.IsLeaf = true
		}
		return node
	}

	topLeftGrid := make([][]int, 0)
	topRightGrid := make([][]int, 0)
	bottomLeftGrid := make([][]int, 0)
	bottomRightGrid := make([][]int, 0)

	for i := 0; i < len(grid); i++ {
		topLeftItem := make([]int, 0, 4)
		topRightItem := make([]int, 0, 4)
		bottomLeftItem := make([]int, 0, 4)
		bottomRightItem := make([]int, 0, 4)

		for j := 0; j < len(grid[i]); j++ {
			if i >= 0 && i <= len(grid)/2 {
				if j >= 0 && j <= len(grid[i])/2 {
					topLeftItem = append(topLeftItem, grid[i][j])
				} else {
					topRightItem = append(topRightItem, grid[i][j])
				}
			} else {
				if j >= 0 && j <= len(grid[i])/2 {
					bottomLeftItem = append(bottomLeftItem, grid[i][j])
				} else {
					bottomRightItem = append(bottomRightItem, grid[i][j])
				}
			}

			if len(topLeftItem) > 0 {
				topLeftGrid = append(topLeftGrid, topLeftItem)
			}
			if len(bottomLeftItem) > 0 {
				bottomLeftGrid = append(bottomLeftGrid, bottomLeftItem)

			}
			if len(topRightItem) > 0 {
				topRightGrid = append(topRightGrid, topRightItem)

			}
			if len(bottomRightItem) > 0 {
				bottomRightGrid = append(bottomRightGrid, bottomRightItem)
			}
		}

	}

	root := &Node{}
	root.TopLeft = construct(topLeftGrid)
	root.TopRight = construct(topRightGrid)
	root.BottomLeft = construct(bottomLeftGrid)
	root.BottomRight = construct(bottomRightGrid)
	return root
}
