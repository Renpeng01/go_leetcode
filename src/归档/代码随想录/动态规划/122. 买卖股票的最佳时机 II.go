package main

func maxProfit(prices []int) int {
	dp := make([][]int, 2)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(prices))
	}

	dp[0][0] = 0
	dp[1][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1]+prices[i])
		dp[1][i] = max(dp[0][i-1]-prices[i], dp[1][i-1])
	}

	return dp[0][len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
