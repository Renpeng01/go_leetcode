package main

import "math"

func minCostClimbingStairs(cost []int) int {

	if len(cost) == 0 {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	if len(cost) == 2 {
		return cost[1]
	}

	dp := make([]int, len(cost)+1)
	dp[1] = cost[0]
	dp[2] = cost[1]

	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1])
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
