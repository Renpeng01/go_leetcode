package main

// 可以转化为01背包问题，物品是否可以把容量为sum一半的背包装满
func canPartition1(nums []int) bool { // 二维dp数组

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	bag := sum / 2

	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bag+1)
	}

	for j := nums[0]; j <= bag; j++ {
		dp[0][j] = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= bag; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}

		}
	}

	return dp[len(nums)-1][bag] == bag
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canPartition(nums []int) bool { // 一维dp数组

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	bag := sum / 2

	dp := make([]int, bag+1)

	for j := nums[0]; j <= bag; j++ {
		dp[j] = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}

	}

	return dp[bag] == bag
}
