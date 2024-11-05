package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
	Pre  *ListNode
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
	list.head.Next = list.tail
	list.tail.Pre = list.head
	return list
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.len {
		return -1
	}
	i := 0
	cur := this.head.Next
	for cur != this.tail {
		if i == index {
			return cur.Val
		}
		i++
		cur = cur.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &ListNode{
		Val:  val,
		Next: this.head.Next,
		Pre:  this.head,
	}
	this.head.Next = newHead
	this.tail.Pre = newHead
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	tail := this.tail

	node := &ListNode{
		Val:  val,
		Next: tail,
		Pre:  this.tail.Pre,
	}
	this.tail.Pre.Next = node
	this.tail.Pre = node
	this.len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= this.len {
		fmt.Println("AddAtIndex error, index >= this.len index: %+v, len: %+v", index, this.len)
		return
	}

	newNode := &ListNode{
		Val: val,
	}
	i := 0
	cur := this.head.Next
	for cur != this.tail {
		if i == index {
			cur.Pre.Next = newNode
			newNode.Pre = cur.Pre
			newNode.Next = cur
			cur.Pre = newNode
			this.len++
			break
		}
		i++
		cur = cur.Next
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.len {
		return
	}
	i := 0
	cur := this.head.Next
	for cur != this.tail {
		if i == index {
			cur.Pre.Next = cur.Next
			cur.Next.Pre = cur.Pre
			cur.Next = nil
			cur.Pre = nil
			this.len--
			break
		}
		i++
		cur = cur.Next
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
	cur := list.head.Next
	res := make([]string, 0, 16)
	for cur != list.tail {
		res = append(res, strconv.Itoa(cur.Val))
		cur = cur.Next
	}
	strs := strings.Join(res, ",")
	fmt.Println(strs)
}
