package main

func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1

	// for i := 0; i <= len(nums); i++ {
	// 	dp[i] = 1
	// }

	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}

		}
	}

	return dp[target]
}
