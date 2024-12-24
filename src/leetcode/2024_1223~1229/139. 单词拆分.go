package main

func wordBreak(s string, wordDict []string) bool {
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

	// fmt.Println(dp)
	return dp[len(s)]
}

func main() {
	s := "leetcode"
	arr := []string{"leet", "code"}
	wordBreak(s, arr)
}
