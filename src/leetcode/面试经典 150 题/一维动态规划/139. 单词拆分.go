package main

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s))

	dict := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		dict[w] = true
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if !dp[i] {
				dp[i] = dict[s[0:i+1]] || (dp[j] && dict[s[j+1:i+1]])
			}
		}
	}
	return dp[len(s)-1]
}
