package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	existedNode := make(map[int]*Node, 16)
	existedEdges := make(map[[2]int]struct{}, 16)

	newNode := &Node{}
	cloneNode(node, newNode, existedNode, existedEdges)
	return newNode
}

func cloneNode(source, d *Node, existedNode map[int]*Node, existedEdges map[[2]int]struct{}) {
	d.Val = source.Val
	d.Neighbors = make([]*Node, 0, len(source.Neighbors))
	existedNode[d.Val] = d

	existedEdgesCnt := 0
	for _, n := range source.Neighbors {
		if _, ok := existedEdges[[2]int{source.Val, n.Val}]; ok {
			existedEdgesCnt++
			continue
		}

		if node, ok := existedNode[n.Val]; ok {
			d.Neighbors = append(d.Neighbors, node)
			existedEdges[[2]int{source.Val, n.Val}] = struct{}{}
		} else {
			newNode := &Node{}
			cloneNode(n, newNode, existedNode, existedEdges)
			d.Neighbors = append(d.Neighbors, newNode)
			existedEdges[[2]int{source.Val, newNode.Val}] = struct{}{}
			existedNode[newNode.Val] = newNode
		}
	}

	if existedEdgesCnt == len(source.Neighbors) {
		return
	}
}
