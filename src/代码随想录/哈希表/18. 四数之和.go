package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 8)

	for j := 0; j < len(nums); j++ {
		if nums[j] > target && nums[j] >= 0 {
			return res
		}

		if j > 0 && nums[j] == nums[j-1] {
			continue
		}

		for i := j + 1; i < len(nums); i++ {
			if nums[j]+nums[i] > target && nums[j]+nums[i] >= 0 {
				break
			}
			if i > j+1 && nums[i] == nums[i-1] {
				continue
			}
			left := i + 1
			right := len(nums) - 1
			for left < right {
				if nums[j]+nums[i]+nums[left]+nums[right] > target {
					right--
				} else if nums[j]+nums[i]+nums[left]+nums[right] < target {
					left++
				} else {
					res = append(res, []int{nums[j], nums[i], nums[left], nums[right]})
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					left++
					right--
				}
			}
		}

	}
	return res
}
