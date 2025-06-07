package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt16
	}
	dp[0] = 0

	for i := 0; i < len(coins); i++ { // 物品
		for j := coins[i]; j <= amount; j++ { // 背包
			dp[j] = min(dp[j-coins[i]]+1, dp[j])
		}
	}

	if dp[amount] == math.MaxInt16 {
		return -1
	}
	return dp[amount]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
