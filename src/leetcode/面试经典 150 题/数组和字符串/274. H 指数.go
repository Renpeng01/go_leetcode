package main

import "sort"

func hIndex(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	res := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] > i {
			res = i + 1
		}
	}
	return res
}

func hIndex1(citations []int) int {

	count := make([]int, len(citations)+1)

	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations) {
			count[len(citations)]++
		} else {
			count[citations[i]]++
		}
	}

	total := 0
	for i := len(count) - 1; i >= 0; i-- {
		total += count[i]
		if total >= i {
			return i
		}

	}
	return 0
}
