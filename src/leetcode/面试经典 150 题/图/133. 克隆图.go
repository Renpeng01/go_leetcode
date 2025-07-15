package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	exist := make(map[[2]int]struct{}, 16)
	new := &Node{}
	cloneExe(node, new, exist)
	return new

}

func cloneExe(source, new *Node, exist map[[2]int]struct{}) {
	new.Val = source.Val
	new.Neighbors = make([]*Node, 0, len(source.Neighbors))

	pathExist := 0
	for _, node := range source.Neighbors {
		if _, ok := exist[[2]int{source.Val, node.Val}]; ok {
			pathExist++
		} else {
			newNode := &Node{}
			cloneExe(node, newNode, exist)
			new.Neighbors = append(new.Neighbors, newNode)
			exist[[2]int{source.Val, node.Val}] = struct{}{}
		}
	}

	if pathExist == len(source.Neighbors) {
		return
	}
}

// func clone(source, new *Node, exist map[[2]int]struct{}) {
// 	new.Val = source.Val
// 	new.Neighbors = make([]*Node, 0, len(source.Neighbors))

// 	pathExist := 0
// 	for _, n := range source.Neighbors {
// 		if _,ok := exist[[2]int[source.Val,n.Val]];ok{
// 			pathExist++
// 		}else{
// exist[[2]int[source.Val,n.Val]] = struct{}{}
// // 		}
// 	}
// }

// func main() {
// 	node1 := &Node{
// 		Val: 1,
// 	}
// 	node2 := &Node{
// 		Val: 2,
// 	}
// 	node3 := &Node{
// 		Val: 3,
// 	}
// 	node4 := &Node{
// 		Val: 4,
// 	}

// 	node1.Neighbors = []*Node{node2, node4}
// 	node2.Neighbors = []*Node{node1, node3}
// 	node3.Neighbors = []*Node{node2, node4}
// 	node4.Neighbors = []*Node{node1, node3}

// 	newNode1 := cloneGraph(node1)

// 	fmt.Println

// }
