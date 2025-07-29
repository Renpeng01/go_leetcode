package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	maxVal := 1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if dp[i] > maxVal {
			maxVal = dp[i]
		}

	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
