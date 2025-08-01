package main

func minCostClimbingStairs(cost []int) int {

	if len(cost) == 0 {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
