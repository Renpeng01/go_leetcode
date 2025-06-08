package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	uniqueCnt := 1

	val := nums[0]
	currentIndex := 1
	for i := 1; i < len(nums); {
		idx := findUniqueIdx(nums, i, len(nums), val)
		if idx == -1 {
			break
		}
		uniqueCnt++
		val = nums[idx]
		nums[currentIndex], nums[idx] = nums[idx], nums[currentIndex]
		i = idx + 1
		currentIndex++

	}

	return uniqueCnt
}

func findUniqueIdx(nums []int, start, end, val int) int {
	res := -1
	for i := start; i < end; i++ {
		if nums[i] != val {
			res = i
			break
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}

	res := removeDuplicates(nums)

	fmt.Println(res)
}
