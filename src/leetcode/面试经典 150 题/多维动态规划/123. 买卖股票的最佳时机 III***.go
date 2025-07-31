package main

func maxProfit(prices []int) int {

	dp := make([][]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}

	// dp[i][0] 不操作
	// dp[i][1] 第一次持有
	// dp[i][2] 第一次不持有
	// dp[i][3] 第二次持有
	// dp[i][4] 第二次不持有

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[len(prices)-1][4]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
