package main

type item struct {
	sum  int
	pair []int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	heap := make([]*item, 0, k+1)
	for i := 0; i < len(nums1); i++ {
		for j := 0; i < len(nums2); j++ {
			it := &item{
				sum:  nums1[i] + nums2[j],
				pair: []int{nums1[i], nums2[j]},
			}
			buildMaxHeap(it, heap, k)
		}
	}

	res := make([][]int, 0, k)

	for _, item := range heap {
		res = append(res, item.pair)
	}
	return res
}

func buildMaxHeap(it *item, heap []*item, k int) {
	if len(heap)-1 < k {
		heap = append(heap, it)

		i := len(heap) - 1
		for i > 1 {
			if i/2 < 1 {
				return
			}

			if heap[i].sum > heap[i/2].sum {
				heap[i], heap[i/2] = heap[i/2], heap[i]
			}
			i = i / 2
		}
	}

	if heap[1].sum < it.sum {
		return
	}
	heap[1] = it

	i := 1

	for i < len(heap) {
		if i*2 >= len(heap) {
			break
		}

		if i*2+1 >= len(heap) {
			if heap[i*2+1].sum > heap[i].sum {
				heap[i*2+1], heap[i] = heap[i], heap[i*2+1]
				break
			}
		}

		swapIndex := i * 2
		if heap[i*2+1].sum > heap[i*2].sum {
			swapIndex = i*2 + 1
		}
		heap[swapIndex], heap[i] = heap[i], heap[swapIndex]
		i = swapIndex
	}

}
