package main

import (
	"fmt"
)

func jump(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := make([][]int, len(nums))
	steps := make([][]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(nums); i++ {
		steps[i] = []int{100, 100}

	}

	dp[0][0] = 0
	dp[0][1] = nums[0]

	steps[0][1] = 0

	for i := 1; i < len(nums); i++ {

		if dp[i-1][0] >= len(nums)-1 && dp[i-1][1] >= len(nums)-1 {
			fmt.Println(nums)
			fmt.Println(dp)
			fmt.Println(steps)
			return min(steps[i-1][0], steps[i-1][1])
		}

		if dp[i-1][0] >= len(nums)-1 {
			fmt.Println(nums)
			fmt.Println(dp)
			fmt.Println(steps)
			return steps[i-1][0]
		}

		if dp[i-1][1] >= len(nums)-1 {
			fmt.Println(nums)
			fmt.Println(dp)
			fmt.Println(steps)
			return steps[i-1][1]
		}
		if max(dp[i-1][0], dp[i-1][1]) >= i {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
			dp[i][1] = i + nums[i]
		}

		if dp[i-1][0] >= i && dp[i-1][1] >= i {
			steps[i][0] = min(steps[i-1][0], steps[i-1][1])
			steps[i][1] = steps[i][0] + 1
		} else if dp[i-1][0] >= i {
			steps[i][0] = steps[i-1][0]
			steps[i][1] = steps[i][0] + 1
		} else if dp[i-1][1] >= i {
			steps[i][0] = steps[i-1][1]
			steps[i][1] = steps[i][0] + 1
		}
	}

	// minStep := math.MaxInt

	// for i := 0; i < len(steps); i++ {
	// 	minStep = min(steps[i][0], steps[i][1])
	// }
	// return minStep
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{3, 4, 3, 2, 5, 4, 3}

	res := jump(nums)
	fmt.Println(res)
}
