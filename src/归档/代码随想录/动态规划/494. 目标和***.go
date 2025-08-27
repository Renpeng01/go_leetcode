package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if (sum-abs(target))%2 != 0 {
		return 0
	}

	bag := (sum + abs(target)) / 2
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bag+1)
	}

	dp[0][0] = 1
	for j := 1; j <= bag; j++ {
		if j == nums[0] {
			dp[0][j] = 1
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= bag; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for _, v := range dp {
		fmt.Println(v)
	}

	return dp[len(nums)-1][bag]
}

func findTargetSumWays2(nums []int, target int) int {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if abs(target) > sum {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	bag := (sum + target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	for _, v := range dp {
		fmt.Println(v)
	}

	return dp[bag]
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func main() {
	// nums := []int{1, 1, 1, 1, 1}
	// target := 3

	nums := []int{0}
	target := 0

	res := findTargetSumWays2(nums, target)
	fmt.Println(res)
}
