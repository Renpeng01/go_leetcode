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

	// ****** 此处不需要对所有元素排序，可以使用【小】顶堆来优化  ps  小顶堆比大顶队好 https://www.bilibili.com/video/BV1Xg41167Lz?vd_source=70c464e99440c207e9933663bb2e5166&spm_id_from=333.788.videopod.sections
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
