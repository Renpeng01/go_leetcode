package main

func numTrees(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += (dp[j] * dp[i-1-j])
			// fmt.Println("i: ", i, ", dp[i]", dp[i], "j: ", j, ", dp[j]", dp[j], "i-1-j: ", i-1-j, ", dp[i-1-j]", dp[i-1-j])
			// fmt.Println("j: ", j, ", dp[j]", dp[j])
			// fmt.Println("i-1-j: ", i-1-j, ", dp[i-1-j]", dp[i-1-j])
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	numTrees(3)

}
