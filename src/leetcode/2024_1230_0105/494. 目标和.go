package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	if abs(target) > sum {
		return 0
	}
	bagSize := (sum + target) / 2
	dp := make([]int, bagSize+1) // dp[j] 装满背包容量i有dp[j]中方法
	dp[0] = 1                    // 如果不是1，则无法计算
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bagSize]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
