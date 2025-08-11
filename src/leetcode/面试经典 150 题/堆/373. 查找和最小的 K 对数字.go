package main

import (
	"fmt"
)

type item struct {
	sum  int
	pair []int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	heap := make([]*item, 0, k+1)
	heap = append(heap, &item{})
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			it := &item{
				sum:  nums1[i] + nums2[j],
				pair: []int{nums1[i], nums2[j]},
			}
			heap = buildMaxHeap(it, heap, k)
		}
	}

	res := make([][]int, 0, k)

	for i, item := range heap {
		if i == 0 {
			continue
		}
		res = append(res, item.pair)
	}
	return res
}

func buildMaxHeap(it *item, heap []*item, k int) []*item {

	if len(heap)-1 < k {
		heap = append(heap, it)
	} else {
		if heap[1].sum < it.sum {
			return heap
		}

		heap[1] = heap[len(heap)-1]
		heap[len(heap)-1] = it
	}

	i := len(heap) - 1
	for i > 1 {
		if i/2 < 1 {

			return heap
		}

		if heap[i].sum < heap[i/2].sum {
			break
		}
		heap[i], heap[i/2] = heap[i/2], heap[i]
		i = i / 2
	}

	return heap
}

func main() {
	nums1 := []int{-10, -4, 0, 0, 6}
	nums2 := []int{3, 5, 6, 7, 8, 100}
	k := 10
	res := kSmallestPairs(nums1, nums2, k)
	fmt.Println(res)
}

func PrintHeap(heap []*item) {
	// return
	// res := make([]string, 0, 16)
	// for i, h := range heap {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	s := "sum:" + strconv.Itoa(h.sum) + "_pair:" + strconv.Itoa(h.pair[0]) + "," + strconv.Itoa(h.pair[1])
	// 	res = append(res, s)
	// }
	// fmt.Println(strings.Join(res, ","))
}
