package main

import "fmt"

type MedianFinder struct {
	heap *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		heap: NewHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.heap.insert(num)
}

func (this *MedianFinder) FindMedian() float64 {

	if len(this.heap.data)-1 == 1 {
		return float64(this.heap.data[1])
	}

	resList := make([]int, 0, 256)

	l := (len(this.heap.data) - 1) % 2

	n := (len(this.heap.data)-1)/2 + 1
	for i := 1; i <= n; i++ {
		resList = append(resList, this.heap.pop())
	}

	var res float64
	if l == 0 {
		res = float64((resList[len(resList)-1] + resList[len(resList)-2])) / float64(2)
	} else {
		res = float64(resList[len(resList)-1])
	}
	for _, v := range resList {
		this.heap.insert(v)

	}
	return res
}

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	data := make([]int, 0, 256)
	data = append(data, 0)
	return &Heap{
		data: data,
	}
}

func (heap *Heap) insert(v int) {
	heap.data = append(heap.data, v)
	heap.heapifyUp(len(heap.data) - 1)
}

func (heap *Heap) heapifyUp(i int) {
	for i > 1 {
		parent := i / 2
		if heap.data[i] < heap.data[parent] {
			heap.data[i], heap.data[parent] = heap.data[parent], heap.data[i]
		}
		i = parent
	}
}

func (heap *Heap) pop() int {
	res := heap.data[1]
	heap.data[1] = heap.data[len(heap.data)-1]
	heap.data = heap.data[:len(heap.data)-1]
	heap.heapifyDown(1)
	return res
}

func (heap *Heap) heapifyDown(i int) {
	last := len(heap.data) - 1
	for {

		left := 2 * i
		right := 2*i + 1
		targetIndex := i

		if left <= last && heap.data[left] < heap.data[targetIndex] {
			targetIndex = left
		}
		if right <= last && heap.data[right] < heap.data[targetIndex] {
			targetIndex = right
		}

		if targetIndex == i {
			break
		}

		heap.data[i], heap.data[targetIndex] = heap.data[targetIndex], heap.data[i]
		i = targetIndex
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	m := Constructor()

	m.AddNum(-1)
	r := m.FindMedian()
	fmt.Println(r)

	m.AddNum(-2)
	r = m.FindMedian()
	fmt.Println(r)

	m.AddNum(-3)
	r = m.FindMedian()
	fmt.Println(r)

	m.AddNum(-4)
	r = m.FindMedian()
	fmt.Println(r)

	m.AddNum(-5)
	r = m.FindMedian()
	fmt.Println(r)
}
