package main

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, v := range strs {
		zeroNum := 0
		oneNum := 0

		for _, c := range v {
			if c == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}
