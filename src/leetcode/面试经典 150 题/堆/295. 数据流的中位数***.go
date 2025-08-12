package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push 和 Pop 会通过 heap.Interface 的方法被调用
	// 这里只是简单的 append
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push 和 Pop 会通过 heap.Interface 的方法被调用
	// 这里只是简单的 append
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// https://www.bilibili.com/video/BV1VX4y1x7YN/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
type MedianFinder struct {
	maxHeap MaxHeap
	minHeap MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: MaxHeap{},
		minHeap: MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num <= this.maxHeap[0] {
		heap.Push(&this.maxHeap, num)
		if this.maxHeap.Len()-this.minHeap.Len() > 1 {
			v := heap.Pop(&this.maxHeap)
			heap.Push(&this.minHeap, v)
		}
	} else {
		heap.Push(&this.minHeap, num)

		if this.minHeap.Len() > this.maxHeap.Len() {
			v := heap.Pop(&this.minHeap)
			heap.Push(&this.maxHeap, v)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {

	if this.minHeap.Len() == this.maxHeap.Len() {
		return (float64(this.maxHeap[0]) + float64(this.minHeap[0])) / float64(2)
	}
	return float64(this.maxHeap[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
