package main

import "fmt"

func maxSubarraySumCircular1(nums []int) int {
	newNums := make([]int, 0, len(nums)*2)
	dp := make([]int, 0, len(newNums)*2)
	for i := 0; i < 2; i++ {
		for _, v := range nums {
			newNums = append(newNums, v)
			dp = append(dp, v)
		}
	}

	fmt.Println(newNums)

	windowSize := len(nums)
	l := 0
	r := 0
	val := newNums[0]

	sum := newNums[0]
	for i := 1; i < len(newNums); i++ {
		r = i
		sum += newNums[i]
		dp[i] = max(dp[i], dp[i-1]+newNums[i])
		fmt.Println("111dp[i]:", dp[i])
		if r-l+1 <= windowSize {
			if sum <= 0 {
				l = i + 1
				r = i + 1
			}
		} else {
			sum -= newNums[l]
			l++
			for m := l; m < r; m++ {
				if newNums[m] < 0 {
					sum -= newNums[m]
					l++
				} else {
					break
				}
			}
		}
		fmt.Println("222dp[i]:", dp[i])
		val = max(dp[i], val)

	}
	return val
}

// https://www.bilibili.com/video/BV1134y1x73z/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166

func maxSubarraySumCircular(nums []int) int {
	dp := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
		sum += nums[i]
	}
	maxVal := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], dp[i])
		maxVal = max(maxVal, dp[i])
	}

	minVal := 0
	for i := 1; i < len(nums)-1; i++ {
		dp[i] = nums[i] + min(dp[i-1], 0)
		minVal = min(minVal, dp[i])
	}

	return max(maxVal, sum-minVal)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
