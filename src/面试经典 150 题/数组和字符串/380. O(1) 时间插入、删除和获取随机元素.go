package main

type RandomizedSet struct {
	dataMap           map[int]*node
	head              *node
	currentNode       *node
	returnCurrentNode *node
}

type node struct {
	val  int
	next *node
	pre  *node
}

func Constructor() RandomizedSet {
	e := RandomizedSet{}
	e.dataMap = make(map[int]*node, 128)
	e.head = new(node)
	e.currentNode = e.head
	e.returnCurrentNode = e.head
	return e
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		return false
	}
	n := &node{val: val}
	n.pre = this.currentNode
	this.currentNode.next = n
	this.currentNode = n
	this.dataMap[val] = n
	return true

}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dataMap[val]; ok {
		deleteNode := this.dataMap[val]
		delete(this.dataMap, val)
		if deleteNode == this.currentNode {
			this.currentNode = deleteNode.pre
		}

		deleteNode.pre.next = deleteNode.next

		if deleteNode.next != nil {
			deleteNode.next.pre = deleteNode.pre
		}

		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	if this.returnCurrentNode.next != nil {
		this.returnCurrentNode = this.returnCurrentNode.next
	} else {
		this.returnCurrentNode = this.head.next
	}

	return this.returnCurrentNode.val
}
