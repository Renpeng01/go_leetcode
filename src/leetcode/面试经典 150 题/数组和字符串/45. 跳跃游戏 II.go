package main

import "math"

func jump(nums []int) int {

	dp := make([]int, len(nums))

	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
