package main

import "sort"

func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	newNums := make([]int, len(nums))
	index := r
	v1, v2 := 0, 0

	for index >= 0 {
		v1 = nums[r] * nums[r]
		v2 = nums[l] * nums[l]
		if v1 > v2 {
			newNums[index] = v1
			r--
		} else {
			newNums[index] = v2
			l++
		}
		index--
	}

	return newNums
}
