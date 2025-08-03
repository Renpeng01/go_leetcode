package main

import (
	"fmt"
)

func longestPalindrome(s string) string {

	dp := make([]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if isPalindrome(s[j : i+1]) {
				dp[i] = max(dp[i-1], i+1-j)
				break
			}
		}
	}

	longest := 0
	endIndex := 0
	for i, v := range dp {
		if v > longest {
			longest = v
			endIndex = i
		}
	}

	return s[endIndex-longest+1 : endIndex+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func main() {
	s := "babad"
	res := longestPalindrome(s)
	fmt.Println(res)
}
