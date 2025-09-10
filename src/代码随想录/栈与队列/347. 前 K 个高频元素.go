package main

import "sort"

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}

	type sortItem struct {
		k   int
		val int
	}

	l := make([]*sortItem, 0, len(m))
	for k, v := range m {
		l = append(l, &sortItem{
			k:   k,
			val: v,
		})
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i].val > l[j].val
	})

	n := k
	if len(l) < k {
		n = len(l)
	}
	res := make([]int, 0, n)

	for _, item := range l[:n] {
		res = append(res, item.k)
	}
	return res

}
