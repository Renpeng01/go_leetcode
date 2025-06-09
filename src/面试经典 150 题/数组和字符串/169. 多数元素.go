package main

import "sort"

func majorityElement1(nums []int) int {

	sort.Ints(nums)

	repeatCnt := 1
	maxReapeatVal := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			repeatCnt++
		} else {
			if repeatCnt > len(nums)/2 {
				return maxReapeatVal
			}
			maxReapeatVal = nums[i]
			repeatCnt = 1
		}
	}

	return maxReapeatVal
}

// 根据题意描述，下标为n/2元素一定是众数
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2] // 中位数
}
