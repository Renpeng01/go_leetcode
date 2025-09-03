package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	n    int
}

func Constructor() MyLinkedList {
	head := new(Node)
	return MyLinkedList{
		head: head,
		n:    0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if !this.checkIndex(index) {
		return -1
	}
	res := this.head
	for i := 0; i <= index; i++ {
		res = res.next

	}
	return res.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)

}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.n-1, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// if !this.checkIndex(index) {
	// 	return
	// }

	if index < 0 || index > this.n {
		return
	}
	pre := this.head
	for i := 0; i < index; i++ {
		pre = pre.next
	}

	newNode := &Node{
		val: val,
	}

	newNode.next = pre.next
	pre.next = newNode
	this.n++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if !this.checkIndex(index) {
		return
	}

	pre := this.head
	for i := 0; i < index; i++ {
		pre = pre.next
	}

	pre.next = pre.next.next
	this.n--
}

func (this *MyLinkedList) checkIndex(index int) bool {
	return index >= 0 && index < this.n
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
// [[],[1],[3],[1,2],[1],[1],[1]]
func main() {

	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)

	fmt.Println("print list")

	head := list.head
	for head.next != nil {
		fmt.Println(head.next.val)
		head = head.next
	}

	get1 := list.Get(1)
	fmt.Println("get1: ", get1)

	list.DeleteAtIndex(1)
	fmt.Println("after delete print list")

	head = list.head
	for head.next != nil {
		fmt.Println(head.next.val)
		head = head.next
	}
	get1 = list.Get(1)
	fmt.Println("get1: ", get1)
}
