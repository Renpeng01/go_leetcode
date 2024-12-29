package main

func fib(n int) int {

	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1) // dp[i] 第i个数为dp[i]
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]

}
