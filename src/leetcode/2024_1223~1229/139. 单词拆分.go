package main

func wordBreak1(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i-len(wordDict[j]) >= 0 {
				if (dp[i-len(wordDict[j])] && s[(i-len(wordDict[j])):i] == wordDict[j]) || s[:i] == wordDict[j] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	set := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		set[v] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && set[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "leetcode"
	arr := []string{"leet", "code"}
	wordBreak(s, arr)
}
