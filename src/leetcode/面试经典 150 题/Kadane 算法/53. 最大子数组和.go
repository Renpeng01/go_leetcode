package main

func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
	}
	val := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], dp[i])
		val = max(val, dp[i])
	}
	return val

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
