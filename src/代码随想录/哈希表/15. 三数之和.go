package main

import "sort"

// hash表
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	m := make(map[int]int, 16)
	for i, v := range nums {
		m[v] = i
	}
	res := make([][]int, 0, 8)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if index, ok := m[0-nums[i]-nums[j]]; ok && index > j {
				res = append(res, []int{nums[i], nums[j], nums[index]})
			}
		}
	}
	return res
}

// 双指针
func threeSum(nums []int) [][]int {

}
