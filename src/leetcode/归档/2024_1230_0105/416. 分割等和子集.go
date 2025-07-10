package main

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = Max(dp[j], dp[j-nums[i]]+nums[i])
			if dp[j] == target {
				return true
			}
		}
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
