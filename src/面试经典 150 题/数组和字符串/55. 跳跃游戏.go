package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1][0] >= len(nums)-1 || dp[i-1][1] >= len(nums)-1 {
			return true
		}
		if max(dp[i-1][0], dp[i-1][1]) >= i {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = i + nums[i]
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{0}

	res := canJump(nums)

	fmt.Println(res)
}
