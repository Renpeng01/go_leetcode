package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt

	sum := 0

	left := 0
	right := 0

	for left <= right && right < len(nums) {
		sum += nums[right]
		if sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			for left < right {
				if (sum - nums[left]) < target {
					break
				}
				if right-(left+1)+1 < minLen {
					minLen = right - (left + 1) + 1
				}
				sum -= nums[left]
				left++
			}
		}
		right++
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
