package main

import "fmt"

type item struct {
	sum  int
	pair []int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	heap := make([]*item, 0, k+1)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
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
	//  len(heap)-1 的值即是长度，也是index
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
		return
	}

	if heap[1].sum < it.sum {
		return
	}
	heap[1] = it

	i := 1

	for i < len(heap) {
		if i*2 >= len(heap)-1 {
			break
		}

		if i*2+1 >= len(heap)-1 {
			if heap[i*2].sum > heap[i].sum {
				heap[i*2], heap[i] = heap[i], heap[i*2]
			}
			break
		}

		swapIndex := i * 2
		if heap[i*2+1].sum > heap[i*2].sum {
			swapIndex = i*2 + 1
		}
		heap[swapIndex], heap[i] = heap[i], heap[swapIndex]
		i = swapIndex
	}
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	res := kSmallestPairs(nums1, nums2, k)
	fmt.Println(res)
}
