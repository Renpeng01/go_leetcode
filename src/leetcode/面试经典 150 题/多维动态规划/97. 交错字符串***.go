package main

// https://www.bilibili.com/video/BV11i4y1D7gj/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
func isInterleave(s1 string, s2 string, s3 string) bool {

	s1Len := len(s1)
	s2Len := len(s2)
	s3Len := len(s3)

	if s1Len == 0 && s2Len == 0 && s3Len == 0 {
		return true
	}

	if s1Len+s2Len != s3Len {
		return false
	}

	// dp[i][j]表示s1的0到i 和 s2的0到j, 是否能够组成s3的0到 i+j
	dp := make([][]bool, s1Len+1)
	for i := 0; i < s1Len; i++ {
		dp[i] = make([]bool, s2Len+1)
	}

	dp[0][0] = true

	for i := 0; i < s1Len; i++ {
		for j := 0; j < s2Len; j++ {
			// if s3[i+j] == s1[i] || s3[i+j] == s2[j] {
			// 	dp[i][j] = dp[i-1][j] || dp[i][j-1]
			// }
			// if s3[i+j] == s1[i] {
			// 	dp[i][j] = dp[i-1][j]
			// }
			// if s3[i+j] == s2[j] {
			// 	dp[i][j] = dp[i][j-1] || dp[i][j]
			// }
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}

		}
	}
	return dp[s1Len-1][s2Len-1]
}
