package main

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt
		for j := 1; j*j <= i; j++ {
			min = int(math.Min(float64(min), float64(dp[i-j*j])))
		}
		dp[i] = min + 1
	}
	return dp[n]
}
