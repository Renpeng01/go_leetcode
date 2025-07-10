package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dpRob1 := make([]int, len(nums))
	dpRob1[0] = nums[0]
	dpRob1[1] = nums[0]

	dpNoRob1 := make([]int, len(nums))
	dpNoRob1[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		if i == len(nums)-1 {
			dpRob1[i] = dpRob1[i-1]
		} else {
			dpRob1[i] = max(dpRob1[i-1], dpRob1[i-2]+nums[i])
		}
		dpNoRob1[i] = max(dpNoRob1[i-1], dpNoRob1[i-2]+nums[i])

	}

	if dpNoRob1[len(nums)-1] > dpRob1[len(nums)-1] {
		return dpNoRob1[len(nums)-1]
	}
	return dpRob1[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
