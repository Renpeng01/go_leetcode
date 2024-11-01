package main

import "sort"

func sortedSquares1(nums []int) []int {
	newNums := make([]int, 0, len(nums))
	for _, v := range nums {
		newNums = append(newNums, v*v)
	}

	sort.Ints(newNums)
	return newNums
}

// 双指针解法
func sortedSquares(nums []int) []int {
	negIndex := -1
	for i, v := range nums {
		if v < 0 {
			negIndex = i
		}
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	if negIndex == -1 {
		return nums
	}

	newNums := make([]int, 0, len(nums))

	aStart := negIndex
	aEnd := 0

	bStart := negIndex + 1
	bEnd := len(nums) - 1

	for aStart >= aEnd && bStart <= bEnd {
		if nums[aStart] < nums[bStart] {
			newNums = append(newNums, nums[aStart])
			aStart--
		} else {
			newNums = append(newNums, nums[bStart])
			bStart++
		}
	}

	for bStart <= bEnd {
		newNums = append(newNums, nums[bStart])
		bStart++
	}

	for aStart >= aEnd {
		newNums = append(newNums, nums[aStart])
		aStart--
	}
	return newNums
}
