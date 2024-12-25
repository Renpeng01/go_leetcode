package main

func maxProduct(nums []int) int {
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))

	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	maxVal := dpMax[0]
	for i := 1; i < len(nums); i++ {
		dpMax[i] = max(dpMax[i-1]*nums[i], max(nums[i], dpMin[i-1]*nums[i]))
		dpMin[i] = min(dpMax[i-1]*nums[i], min(nums[i], dpMin[i-1]*nums[i]))

		if dpMax[i] > maxVal {
			maxVal = dpMax[i]
		}
	}
	return maxVal
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
