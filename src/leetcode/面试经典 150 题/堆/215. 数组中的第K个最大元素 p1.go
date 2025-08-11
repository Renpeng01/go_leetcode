package main

var heap []int

func findKthLargest(nums []int, k int) int {

	heap = make([]int, 0, len(nums)+1)
	heap = append(heap, 0)

	for _, v := range nums {
		buildMinHeap(v, k)
	}

	return heap[1]
}

func buildMinHeap(v int, k int) {
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
