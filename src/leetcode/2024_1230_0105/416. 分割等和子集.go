package main

import "fmt"

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	dp := make([]int, sum/2+1)
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= nums[i]; j-- {
			dp[j] = Max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	fmt.Println(dp)

	return dp[sum/2] == sum/2

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	canPartition([]int{1, 5, 11, 5})
}
