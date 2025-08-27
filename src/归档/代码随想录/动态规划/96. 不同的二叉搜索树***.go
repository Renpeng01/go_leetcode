package main

// https://www.bilibili.com/video/BV1eK411o7QA?spm_id_from=333.788.videopod.sections&vd_source=70c464e99440c207e9933663bb2e5166
func numTrees(n int) int {

	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
