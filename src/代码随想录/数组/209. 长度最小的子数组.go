package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := nums[0]
	if sum >= target {
		return 1
	}

	minLen := math.MaxInt
	l := 0
	for r := 1; r <= len(nums)-1; r++ {
		sum += nums[r]
		if sum >= target {
			if r-l+1 < minLen {
				minLen = r - l + 1
			}

			tmpLen := r - l + 1
			for l < r {
				sum -= nums[l]
				if sum < target {
					break
				}
				l++
				tmpLen--
			}

			if tmpLen < minLen {
				minLen = tmpLen
			}
		}
	}
	return minLen
}
