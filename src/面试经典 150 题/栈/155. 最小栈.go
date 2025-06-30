package main

import "sort"

type MinStack struct {
	data       []int
	sortedData []int
}

func Constructor() MinStack {
	data := make([]int, 0, 256)
	sortedData := make([]int, 0, 256)
	stack := MinStack{
		data:       data,
		sortedData: sortedData,
	}
	return stack
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	this.sortedData = append(this.sortedData, val)
	sort.Ints(this.sortedData)
}

func (this *MinStack) Pop() {
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

func (this *MinStack) Top() int {
	if this != nil || len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	return this.sortedData[0]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
