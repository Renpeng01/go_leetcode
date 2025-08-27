package main

func isSubsequence(s string, t string) bool {

	if len(s) == 0 && len(t) > 0 {
		return true
	}

	if len(s) > 0 && len(t) == 0 {
		return false
	}

	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
