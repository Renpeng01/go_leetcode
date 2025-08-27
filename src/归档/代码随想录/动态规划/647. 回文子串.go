package main

// 相当于暴力解法
func countSubstrings1(s string) int {

	dp := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		dp[i] = 1
	}

	sum := dp[0]
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			// 没有递推关系，属于暴力算法 *******
			// if isPalindrome(s[j : i+1]) {
			// 	dp[i]++ //
			// }

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

func countSubstrings(s string) int {
	// dp[i][j] [i,j]之间的子串是否是回文
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					res++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res

}
