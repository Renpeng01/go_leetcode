package main

import (
	"fmt"
	"sort"
)

type Project struct {
	Profit  int
	Capital int
}

var heap []int

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	heap = make([]int, 0, k+1)
	heap = append(heap, 0)

	projects := make([]*Project, 0, len(profits))
	for i := 0; i < len(profits); i++ {
		projects = append(projects, &Project{
			Profit:  profits[i],
			Capital: capital[i],
		})
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].Capital < projects[j].Capital
	})

	maxHeap := NewHeap()
	visited := make(map[int]struct{}, len(projects))
	for k > 0 {
		for i, v := range projects {
			if _, ok := visited[i]; ok {
				continue
			}
			if w < v.Capital {
				break
			}
			maxHeap.insert(v.Profit)
			visited[i] = struct{}{}
		}

		if len(maxHeap.data) <= 1 {
			break
		}
		w = w + maxHeap.pop()
		k--
	}

	return w
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

	i = i / 2
	for i > 1 {
		parent := i / 2
		if heap.data[i] > heap.data[parent] {
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

		if left <= last && heap.data[left] > heap.data[targetIndex] {
			targetIndex = left
		}
		if right <= last && heap.data[right] > heap.data[targetIndex] {
			targetIndex = right
		}

		if targetIndex == i {
			break
		}

		heap.data[i], heap.data[targetIndex] = heap.data[targetIndex], heap.data[i]
		i = targetIndex
	}
}

func main() {
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}

	k := 2
	w := 0

	res := findMaximizedCapital(k, w, profits, capital)
	fmt.Println(res)

}
