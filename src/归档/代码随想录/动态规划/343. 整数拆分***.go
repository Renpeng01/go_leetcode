package main

// 本题难点在于想不到用动态规划的方法解决
func integerBreak(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
			dp[i] = max(dp[i], max(dp[j]*(i-j), j*(i-j))) // 这2个都可以
		}
	}

	return dp[n]

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
