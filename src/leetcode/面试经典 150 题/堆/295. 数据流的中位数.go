package main

type MedianFinder struct {
	data []int
}

func Constructor() MedianFinder {
	data := make([]int, 0, 256)
	return MedianFinder{
		data: data,
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.data = append(this.data, num)

}

func (this *MedianFinder) FindMedian() float64 {

	// len为奇数 len/2 为中位数
	// len为偶数 (len/2-1 + len/2）/2 为中位数

	n := len(this.data) / 2
	mHeap := NewHeap(n)
	for i := 0; i < len(this.data); i++ {
		mHeap.insert(this.data[i])
	}

	if len(this.data)%2 == 1 {
		return float64(mHeap.pop())
	}

	return (float64(mHeap.pop()) + float64(mHeap.pop())) / 2

}

type Heap struct {
	data []int
	k    int
}

func NewHeap(k int) *Heap {
	data := make([]int, 0, 256)
	data = append(data, 0)
	return &Heap{
		data: data,
		k:    k,
	}
}

func (heap *Heap) insert(v int) {

	if len(heap.data)-1 < heap.k {
		heap.data = append(heap.data, v)
		heap.heapifyUp(len(heap.data) - 1)
	} else {
		if heap.data[1] > v {
			return
		}
		heap.data[1] = v
		heap.heapifyDown(1)
	}
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
