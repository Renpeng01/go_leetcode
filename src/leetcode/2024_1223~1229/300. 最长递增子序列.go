package main

import (
	"math"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	max := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}

	return max
}

func main() {
	lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
}
