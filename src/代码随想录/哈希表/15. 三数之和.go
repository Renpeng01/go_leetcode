package main

import "sort"

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	m := make(map[int]int, 16)
	for i, v := range nums {
		m[v] = i
	}

	existed := make(map[[3]int]struct{})

	res := make([][]int, 0, 8)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if index, ok := m[0-nums[i]-nums[j]]; ok && index > j {
				if _, existedOk := existed[[3]int{nums[i], nums[j], nums[index]}]; !existedOk {
					res = append(res, []int{nums[i], nums[j], nums[index]})
					existed[[3]int{nums[i], nums[j], nums[index]}] = struct{}{}
				}
			}
		}
	}
	return res
}
