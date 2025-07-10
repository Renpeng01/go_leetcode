package main

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	// 0 不持有
	// 1 持有
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(prices[i]+dp[i-1][1], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		// dp[i][1] = max(-prices[i], dp[i-1][1])  区别是 dp[i-1][0]-prices[i]（多次购买，当前手头可能有钱） 和 -prices[i]
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
