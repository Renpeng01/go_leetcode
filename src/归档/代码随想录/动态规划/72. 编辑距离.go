package main

func minDistance(word1 string, word2 string) int {
	// dp[i][j] 以i-1为结尾的word1 以j-1为结尾的word2 相同的最少操作次数
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
