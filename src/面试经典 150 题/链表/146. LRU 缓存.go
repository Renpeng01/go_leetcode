package main

import "fmt"

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
		m:        make(map[int]*listNode, capacity),
		capacity: capacity,
		tail:     head,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}

	if this.head.next != node {

		if this.tail == node {
			this.tail = this.tail.pre

			node.pre.next = nil

			node.next = this.head.next
			this.head.next.pre = node
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
	return node.val
}

func (this *LRUCache) Put(key int, value int) {

	node, ok := this.m[key]
	if !ok {
		newNode := &listNode{
			val: value,
			key: key,
		}
		if this.head.next == nil {
			this.head.next = newNode
			newNode.pre = this.head
			this.tail = newNode
		} else {
			newNode.next = this.head.next
			this.head.next.pre = newNode
			this.head.next = newNode
			newNode.pre = this.head
		}
		if len(this.m) >= this.capacity {
			delete(this.m, this.tail.key)
		}
		this.m[key] = newNode
	} else {
		if node == this.head.next {
			return
		}

		if node == this.tail {
			this.tail = node.pre
		}
		node.pre.next = node.next
		node.next.pre = node.pre

		node.next = this.head.next
		this.head.next.pre = node
		node.pre = this.head
		this.head.next = node

	}
}

func main() {

	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	// head := lru.head
	// tail := lru.tail

	// fmt.Println("head: ", head.key, head.val)
	// fmt.Println("tail: ", tail.key, tail.val)

	// cur := head.next
	// for cur != nil {
	// 	fmt.Println(cur.key, cur.val)
	// 	cur = cur.next
	// }

	// fmt.Println("m: ", lru.m)
	// fmt.Println("capacity: ", lru.capacity)

	res := lru.Get(1)
	fmt.Println("get(1):", res)
	lru.Put(3, 3)
	res = lru.Get(2)
	fmt.Println("get(2):", res)
	// lru.Put(4, 4)
	// res = lru.Get(1)
	// fmt.Println("get(1):", res)
	// res = lru.Get(3)
	// fmt.Println("get(3):", res)
	// res = lru.Get(4)
	// fmt.Println("get(4):", res)
}

// [null,null,null,1,null,-1,null,1,3,4]
// 预期结果
// [null,null,null,1,null,-1,null,-1,3,4]
