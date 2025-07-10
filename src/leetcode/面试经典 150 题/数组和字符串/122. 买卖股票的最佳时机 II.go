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
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 注意这里与121的区别，这里可以多次持有并卖出
	}

	return dp[len(prices)-1][0]

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit1(prices []int) int {
	min := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > 0 {
				maxProfit += prices[i] - min
				if i < len(prices)-1 {
					min = prices[i]
				}
			}
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		maxProfit = max(0, prices[i]-prices[i-1])
	}
	return maxProfit
}
