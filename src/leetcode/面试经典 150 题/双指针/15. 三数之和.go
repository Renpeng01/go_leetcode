package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 8)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left+1 <= len(nums)-1 && nums[left+1] == nums[left] {

					left++
				}
				for right-1 >= 0 && nums[right-1] == nums[right] {

					right--
				}
				left++
				right--
				continue
			}

			if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
