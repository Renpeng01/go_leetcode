package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == q {
		return q
	}
	pPath := getNodePath(root, p)
	qPath := getNodePath(root, q)

	pLen := len(pPath)
	qLen := len(qPath)
	len := qLen
	if qLen > pLen {
		len = pLen
	}

	for i := 1; i < len; i++ {
		if pPath[i] != qPath[i] {
			return pPath[i-1]
		}
	}
	return nil
}

// 先序遍历找到节点路径
func getNodePath(root, target *TreeNode) []*TreeNode {
	path := make([]*TreeNode, 0, 8)
	queue := make([]*TreeNode, 0, 8)

	queue = append(queue, root)
	cur := root

	for cur != nil {
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		if cur == target {
			path = append(path, cur)
			return path
		}

		if cur.Left == nil && cur.Right == nil {
			path = []*TreeNode{}
			cur = queue[0]
			queue = queue[1:]
			continue
		}
		path = append(path, cur)
		cur = cur.Left
	}

	return nil
}

func main() {
	node0 := &TreeNode{
		Val: 3,
	}
	node1 := &TreeNode{
		Val: 5,
	}
	node2 := &TreeNode{
		Val: 1,
	}
	node0.Left = node1
	node0.Right = node2

	node3 := &TreeNode{
		Val: 6,
	}
	node4 := &TreeNode{
		Val: 2,
	}

	node1.Left = node3
	node1.Right = node4

	node5 := &TreeNode{
		Val: 0,
	}
	node6 := &TreeNode{
		Val: 8,
	}

	node2.Left = node5
	node2.Right = node6

	node7 := &TreeNode{
		Val: 7,
	}
	node8 := &TreeNode{
		Val: 4,
	}

	node4.Left = node7
	node4.Right = node8

	path1 := getNodePath(node0, node8)
	for _, v := range path1 {
		fmt.Println(v.Val)
	}
}
