package main

import (
	"fmt"
	"strings"
)

type listNode struct {
	val  int
	key  int
	next *listNode
	pre  *listNode
}

type LRUCache struct {
	head     *listNode
	tail     *listNode
	m        map[int]*listNode
	capacity int
}

func Constructor(capacity int) LRUCache {

	head := &listNode{}
	return LRUCache{
		head:     head,
		tail:     head,
		m:        make(map[int]*listNode, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}
	res := node.val
	this.touch(node)
	return res
}

func (this *LRUCache) touch(node *listNode) {

	if this.head.next == node {
		return
	}
	if this.tail == node {
		this.tail = this.tail.pre
		this.tail.next = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
	node.pre = this.head

}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok {
		node.val = value
		this.touch(node)
	} else {
		if len(this.m) >= this.capacity {
			targetNode := this.tail

			delete(this.m, targetNode.key)
			this.m[key] = targetNode
			targetNode.key = key
			targetNode.val = value
			this.touch(targetNode)

		} else {
			newNode := &listNode{
				key: key,
				val: value,
			}

			if this.head.next == nil {
				this.tail = newNode
				this.head.next = newNode
				newNode.pre = this.head
			} else {
				newNode.next = this.head.next
				this.head.next.pre = newNode
				this.head.next = newNode
				newNode.pre = this.head
			}
			this.m[key] = newNode
		}
	}
}

func main() {
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	// ["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	// [[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]

	var res int

	lru := Constructor(3)
	lru.Put(1, 1)
	PrintLRU(&lru, "Put(1, 1)")

	lru.Put(2, 2)
	PrintLRU(&lru, "Put(2, 2)")

	lru.Put(3, 3)
	PrintLRU(&lru, "Put(3, 3)")

	lru.Put(4, 4)
	PrintLRU(&lru, "Put(4, 4)")

	// res = lru.Get(4)
	// PrintLRU(&lru, "Get(4)")

	// res = lru.Get(3)
	// PrintLRU(&lru, "Get(4)")

	// res = lru.Get(2)
	// PrintLRU(&lru, "Get(4)")

	// res = lru.Get(1)
	// PrintLRU(&lru, "Get(4)")

	// lru.Put(5, 5)
	// PrintLRU(&lru, "Put(5, 5)")

	// res = lru.Get(1)
	// PrintLRU(&lru, "Get(1)")

	// res = lru.Get(2)
	// PrintLRU(&lru, "Get(2)")

	// res = lru.Get(3)
	// PrintLRU(&lru, "Get(3)")

	// res = lru.Get(4)
	// PrintLRU(&lru, "Get(4)")

	// res = lru.Get(5)
	// PrintLRU(&lru, "Get(5)")

	fmt.Println(res)

}

func PrintLRU(lru *LRUCache, tag string) {

	nodeStr := make([]string, 0, 8)
	curNode := lru.head.next
	for curNode != nil {
		nodeStr = append(nodeStr, fmt.Sprintf("%v:%v", curNode.key, curNode.val))
		curNode = curNode.next
	}

	fmt.Printf("%s m: %+v, nodes: %s, fisrt: %+v, tail: %+v\n", tag, lru.m, strings.Join(nodeStr, ","), lru.head.next.key, lru.tail.key)

}
