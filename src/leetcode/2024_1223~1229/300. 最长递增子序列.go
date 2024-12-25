package main

import "math"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
			}
		}
	}

	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
