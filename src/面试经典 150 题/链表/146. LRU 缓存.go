package main

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

				this.tail = this.tail.pre
				this.tail.next = nil
			}
			this.m[key] = newNode
		}
	}
}
