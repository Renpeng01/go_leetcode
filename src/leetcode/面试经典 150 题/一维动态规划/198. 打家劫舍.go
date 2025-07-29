package main

func rob(nums []int) int {
	dp := make([][]int, 2)
	dp[0] = make([]int, len(nums))
	dp[1] = make([]int, len(nums))

	dp[0][0] = 0
	dp[1][0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1])
		dp[1][i] = dp[0][i-1] + nums[i]
	}
	return max(dp[0][len(nums)-1], dp[1][len(nums)-1])

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
