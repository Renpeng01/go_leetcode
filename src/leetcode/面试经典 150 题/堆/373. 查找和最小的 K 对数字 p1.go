package main

import (
	"fmt"
	"strconv"
	"strings"
)

type item struct {
	sum  int
	pair []int
}

var heap []*item

// 效率低下
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	heap = make([]*item, 0, k+1)
	heap = append(heap, &item{})

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if len(heap)-1 == k && nums1[i]+nums2[j] >= heap[1].sum { // 如无此代码，则超时  含义 如果当前sum大于堆顶，则后续也不会出现更小的
				break
			}
			it := &item{
				sum:  nums1[i] + nums2[j],
				pair: []int{nums1[i], nums2[j]},
			}

			buildMaxHeap(it, k)
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

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}

func buildMaxHeap(it *item, k int) {
	if len(heap)-1 < k {
		heap = append(heap, it)
		i := len(heap) - 1

		// heapifyUp
		for i > 1 {
			parent := i / 2
			if heap[i].sum > heap[parent].sum {
				heap[i], heap[parent] = heap[parent], heap[i]
			}
			i = parent
		}
	} else {
		if heap[1].sum <= it.sum {
			// return heap
			return
		}
		i := 1
		heap[i] = it

		// heapifyDown
		last := len(heap) - 1
		for {
			left := 2 * i
			right := 2*i + 1
			maxIndex := i

			if left <= last && heap[left].sum > heap[maxIndex].sum {
				maxIndex = left
			}
			if right <= last && heap[right].sum > heap[maxIndex].sum {
				maxIndex = right
			}

			if maxIndex == i {
				break
			}

			heap[i], heap[maxIndex] = heap[maxIndex], heap[i]
			i = maxIndex
		}
	}
	// return heap
}

func main() {
	nums1 := []int{0, 0, 0, 0, 0}
	nums2 := []int{-3, 22, 35, 56, 76}
	k := 22
	res := kSmallestPairs(nums1, nums2, k)
	fmt.Println(res)
}

func PrintHeap(heap []*item) {
	// return
	res := make([]string, 0, 16)
	for i, h := range heap {
		if i == 0 {
			continue
		}
		s := "sum:" + strconv.Itoa(h.sum) + "_pair:" + strconv.Itoa(h.pair[0]) + "," + strconv.Itoa(h.pair[1])
		res = append(res, s)
	}
	fmt.Println(strings.Join(res, ","))
}

// 性能更好的思路
// https://www.bilibili.com/video/BV1Mv4y1Z79v/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
