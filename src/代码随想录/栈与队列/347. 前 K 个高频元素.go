package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)
	res := make([]int, 0, 16)
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if len(res) == 0 {
			res = append(res, nums[i])
			cnt++
			continue
		}

		if cnt == k {
			break
		}

		if nums[i] == res[len(res)-1] {
			continue
		}
		res = append(res, nums[i])
		cnt++
	}
	return res
}
