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

		node.next = this.tail
		this.tail.pre = node

		this.head.next = node
		node.pre = this.head
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
		node.next = this.head.next
		this.head.next.pre = node
		this.head.next = node
		node.pre = this.head
	}

}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok {
		node.val = value
	} else {
		if len(this.m) >= this.capacity {
			targetNode := this.tail
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

	lru := Constructor(2)
	lru.Put(1, 1)
	PrintLRU(&lru, "Put(1, 1)")

	lru.Put(1, 1)
	PrintLRU(&lru, "Put(2, 2)")

	// res := lru.Get(1)
	// fmt.Printf("Get(1) res: %+v", res)
	// PrintLRU(&lru, "Get(1)")

}

func PrintLRU(lru *LRUCache, tag string) {

	nodeStr := make([]string, 0, 8)
	curNode := lru.head.next
	if curNode != nil {
		nodeStr = append(nodeStr, fmt.Sprintf("%v:%v", curNode.key, curNode.val))
		curNode = curNode.next
	}

	fmt.Printf("%s m: %+v, nodes: %s\n", tag, lru.m, strings.Join(nodeStr, ","))
}
