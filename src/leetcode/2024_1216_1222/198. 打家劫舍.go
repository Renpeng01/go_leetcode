package main

import "math"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	dp[2] = nums[1]
	if dp[1] > dp[2] {
		dp[2] = dp[1]
	}

	for i := 3; i <= len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i-1])))
	}

	return dp[len(nums)]
}
