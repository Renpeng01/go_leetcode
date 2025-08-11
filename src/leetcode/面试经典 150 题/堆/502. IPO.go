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

	for _, v := range projects {

		buildMinHeap(v.Profit, k)
	}

	fmt.Println(heap)
	res := w
	for i := 0; i < len(heap); i++ {
		if i == 0 {
			continue
		}
		res += heap[i]
	}

	return res
}

func buildMinHeap(v, k int) {
	if len(heap)-1 < k {
		heap = append(heap, v)

		i := len(heap) - 1

		// heapifyUp
		for i > 1 {
			parent := i / 2
			if heap[i] < heap[parent] {
				heap[i], heap[parent] = heap[parent], heap[i]
			}
			i = parent
		}

	} else {
		if heap[1] >= v {
			// return heap
			return
		}
		i := 1
		heap[i] = v

		// heapifyDown
		last := len(heap) - 1
		for {
			left := 2 * i
			right := 2*i + 1
			smallIndex := i

			if left <= last && heap[left] < heap[smallIndex] {
				smallIndex = left
			}
			if right <= last && heap[right] < heap[smallIndex] {
				smallIndex = right
			}

			if smallIndex == i {
				break
			}

			heap[i], heap[smallIndex] = heap[smallIndex], heap[i]
			i = smallIndex
		}
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
