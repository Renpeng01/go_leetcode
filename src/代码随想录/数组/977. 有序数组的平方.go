package main

import "sort"

func sortedSquares(nums []int) []int {
	newNums := make([]int, 0, len(nums))
	for _, v := range nums {
		newNums = append(newNums, v*v)
	}

	sort.Ints(newNums)
	return newNums
}
