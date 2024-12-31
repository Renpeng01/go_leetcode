package main

func lastStoneWeightII(stones []int) int {

	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2

	dp := make([]int, target+1)

	for i := 1; i <= target; i++ {
		if i >= stones[0] {
			dp[i] = stones[0]
		}
	}

	for i := 1; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[target]*2
}
