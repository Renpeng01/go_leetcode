package main

func maxProfit(prices []int) int {

	if len(prices) == 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	dp[1][1] = max(dp[0][1], dp[0][0]-prices[1])

	for i := 2; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}

	return max(dp[len(dp)-1][0], dp[len(dp)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
