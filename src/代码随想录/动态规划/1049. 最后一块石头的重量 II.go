package main

func lastStoneWeightII(stones []int) int {

	sum := 0
	for _, v := range stones {
		sum += v
	}

	bag := sum / 2

	dp := make([]int, bag+1)

	for i := stones[0]; i <= bag; i++ {
		dp[i] = stones[0]
	}

	for i := 1; i < len(stones); i++ {
		for j := bag; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[bag]*2

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
