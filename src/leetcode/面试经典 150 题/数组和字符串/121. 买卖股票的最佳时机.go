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
		dp[i][1] = max(dp[i-1][1], -prices[i]) // dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])  这里不能写成这样，因为是题目说智能买一次
	}

	return dp[len(prices)-1][0]

}
func maxProfit1(prices []int) int {
	min := prices[0]

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > maxProfit {
				maxProfit = prices[i] - min
			}
		}
	}

	return maxProfit

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
