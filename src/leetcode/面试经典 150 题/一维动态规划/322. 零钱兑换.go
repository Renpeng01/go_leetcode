package main

import "math"

// 完全背包
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}

		}
	}

	if dp[amount] != math.MaxInt {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
