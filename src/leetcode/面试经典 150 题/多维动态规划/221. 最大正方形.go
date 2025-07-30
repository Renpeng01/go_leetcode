package main

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	dp[0][0] = 0
	if matrix[0][0] == '1' {
		dp[0][0] = 1
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}

	// for _, v := range dp {
	// 	fmt.Println(v)
	// }
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {

			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}

		}
	}

	// for _, v := range dp {
	// 	fmt.Println(v)
	// }

	maxEdge := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] > maxEdge {
				maxEdge = dp[i][j]
			}
		}
	}
	return maxEdge * maxEdge
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	res := maximalSquare(matrix)
	fmt.Println(res)
}
