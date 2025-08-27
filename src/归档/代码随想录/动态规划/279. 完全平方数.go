package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {

	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i <= n/2+1; i++ {
		for j := i * i; j <= n; j++ {
			if dp[j-i*i] != math.MinInt {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}

	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	res := numSquares(11)
	fmt.Println(res)
}
