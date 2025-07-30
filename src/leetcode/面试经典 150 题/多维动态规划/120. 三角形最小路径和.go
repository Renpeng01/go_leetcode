package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]
	dp[1][0] = triangle[0][0] + triangle[1][0]
	dp[1][1] = triangle[0][0] + triangle[1][1]

	for i := 2; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			upLen := len(triangle[i-1]) - 1
			if j == 0 {
				dp[i][j] = dp[i-1][0] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = dp[i-1][upLen] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j]+triangle[i][j], dp[i-1][j-1]+triangle[i][j])
			}

		}
	}

	for _, v := range dp {
		fmt.Println(v)
	}

	minSum := math.MaxInt

	for _, v := range dp[len(dp)-1] {
		if v < minSum {
			minSum = v
		}
	}
	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	res := minimumTotal(triangle)
	fmt.Println(res)

}
