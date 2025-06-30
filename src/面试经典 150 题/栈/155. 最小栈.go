package main

import "sort"

type MinStack1 struct {
	data       []int
	sortedData []int
}

func Constructor1() MinStack1 {
	data := make([]int, 0, 256)
	sortedData := make([]int, 0, 256)
	stack := MinStack1{
		data:       data,
		sortedData: sortedData,
	}
	return stack
}

func (this *MinStack1) Push(val int) {
	this.data = append(this.data, val)
	this.sortedData = append(this.sortedData, val)
	sort.Ints(this.sortedData)
}

func (this *MinStack1) Pop() {
	deleteVal := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	for i := len(this.sortedData) - 1; i >= 0; i-- {
		if this.sortedData[i] == deleteVal {

			if len(this.data) == 0 {
				this.sortedData = make([]int, 0, 256)
				break
			}
			sortedData := make([]int, 0, len(this.data)-1)
			pre := this.sortedData[:i]
			next := this.sortedData[i+1:]
			sortedData = append(sortedData, pre...)
			sortedData = append(sortedData, next...)
			this.sortedData = sortedData
			break
		}
	}
}

func (this *MinStack1) Top() int {
	if this != nil || len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return -1
}

func (this *MinStack1) GetMin() int {
	return this.sortedData[0]

}

/**
 * Your MinStack1 object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type MinStack struct {
	data     []int
	minStack []int
}

func Constructor() MinStack {
	data := make([]int, 0, 256)
	minStack := make([]int, 0, 256)
	stack := MinStack{
		data:     data,
		minStack: minStack,
	}
	return stack
}

func (this *MinStack) Push(val int) {

	this.data = append(this.data, val)
	if len(this.data) == 1 {
		this.minStack = append(this.minStack, val)
		return
	}

	if val < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minStack = this.data[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
