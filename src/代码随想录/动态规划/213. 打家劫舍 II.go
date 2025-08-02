package main

func rob(nums []int) int {

	dp := make([]int, len(nums)+1)

	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {

		dp[i] = max(dp[i], dp[i-2]+nums[i]-)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
