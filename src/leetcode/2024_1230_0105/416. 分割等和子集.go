package main

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2

	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = Max(dp[j], dp[j-nums[i]]+nums[i])
			if dp[j] == target {
				return true
			}
		}
	}
	return false
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
