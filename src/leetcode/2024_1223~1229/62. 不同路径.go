package main

func uniquePaths(m int, n int) int {

	if n <= 0 || m <= 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < n; i++ { // row
		for j := 1; j < m; j++ { // col
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}
