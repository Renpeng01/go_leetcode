package main

import "math"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := math.MinInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}

	}

	if maxProfit < 0 {
		return 0
	}

	return maxProfit
}
