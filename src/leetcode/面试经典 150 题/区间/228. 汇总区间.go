package main

import "strconv"

func summaryRanges(nums []int) []string {

	res := make([]string, 0, len(nums))
	for i := 0; i < len(nums); {
		start := i
		end := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j-1]+1 != nums[j] {

				break
			}
			end = j
		}

		if nums[start] == nums[end] {
			res = append(res, strconv.Itoa(nums[start]))
		} else {
			res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
		}

		i = end + 1
	}
	return res
}
