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
	cur := head
	for i := 0; i < capacity; i++ {
		cur.next = &listNode{}
		cur = cur.next
	}
	return LRUCache{
		head:     head,
		m:        make(map[int]*listNode, capacity),
		capacity: capacity,
		tail:     cur,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
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
		if len(this.m) >= this.capacity {

			newNode.next = this.head.next
			this.head.next.pre = newNode

			this.head.next = newNode
			newNode.pre = this.head

			delete(this.m, this.tail.key)
			this.tail = this.tail.pre
			this.tail.next = nil
		} else {
			this.tail.next = newNode
			this.tail = this.tail.next
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

	res := lru.Get(1)
	fmt.Println("get(1):", res)
	lru.Put(3, 3)
	res = lru.Get(2)
	fmt.Println("get(2):", res)
	lru.Put(4, 4)
	res = lru.Get(1)
	fmt.Println("get(1):", res)
	res = lru.Get(3)
	fmt.Println("get(3):", res)
	res = lru.Get(4)
	fmt.Println("get(4):", res)

}
