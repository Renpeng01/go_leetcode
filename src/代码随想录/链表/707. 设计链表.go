package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	val  int
	next *ListNode
	prev *ListNode
}

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	len  int
}

func Constructor() MyLinkedList {
	list := MyLinkedList{}
	list.head = &ListNode{}
	list.tail = &ListNode{}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.len {
		return -1
	}
	i := 0
	cur := this.head.next
	for cur != this.tail {
		if i == index {
			return cur.val
		}
		i++
		cur = cur.next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &ListNode{
		val:  val,
		next: this.head.next,
		prev: this.head,
	}
	this.head.next = newHead
	this.tail.prev = newHead
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	tail := this.tail

	node := &ListNode{
		val:  val,
		next: tail,
		prev: this.tail.prev,
	}
	this.tail.prev.next = node
	this.tail.prev = node
	this.len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= this.len {
		fmt.Println("AddAtIndex error, index >= this.len index: %+v, len: %+v", index, this.len)
		return
	}

	newNode := &ListNode{
		val: val,
	}
	i := 0
	cur := this.head.next
	for cur != this.tail {
		if i == index {
			cur.prev.next = newNode
			newNode.prev = cur.prev
			newNode.next = cur
			cur.prev = newNode
			this.len++
			break
		}
		i++
		cur = cur.next
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.len {
		return
	}
	i := 0
	cur := this.head.next
	for cur != this.tail {
		if i == index {
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			cur.next = nil
			cur.prev = nil
			this.len--
			break
		}
		i++
		cur = cur.next
	}
}

func (this *MyLinkedList) GetHead() *ListNode {
	return this.head
}

func main() {

	linkList := Constructor()
	linkList.AddAtHead(1)
	// print
	printList(&linkList)

	linkList.AddAtTail(3)
	printList(&linkList)

	linkList.AddAtIndex(1, 2)
	printList(&linkList)

	v := linkList.Get(1)
	fmt.Println("Get(1) ", v)

	linkList.DeleteAtIndex(1)
	printList(&linkList)

	v = linkList.Get(1)
	fmt.Println("Get(1) ", v)

}

func printList(list *MyLinkedList) {
	cur := list.head.next
	res := make([]string, 0, 16)
	for cur != list.tail {
		res = append(res, strconv.Itoa(cur.val))
		cur = cur.next
	}
	strs := strings.Join(res, ",")
	fmt.Println(strs)
}
