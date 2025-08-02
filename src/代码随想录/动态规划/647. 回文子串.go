package main

func countSubstrings(s string) int {

	dp := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		dp[i] = 1
	}

	sum := dp[0]
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if isPalindrome(s[j : i+1]) {
				dp[i]++
			}
		}
		sum += dp[i]
	}

	return sum

}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
