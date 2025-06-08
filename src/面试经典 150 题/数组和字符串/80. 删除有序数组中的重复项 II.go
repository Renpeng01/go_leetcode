package main

func removeDuplicates(nums []int) int {
	slow := 1
	fast := 1

	duplicatesCnt := 0
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
			duplicatesCnt = 0
		} else {
			if duplicatesCnt == 0 {
				nums[slow] = nums[fast]
				slow++
				duplicatesCnt++
			}
		}
		fast++

	}
	return slow
}

func removeDuplicates1(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}
	slow := 2
	fast := 2

	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++

	}
	return slow
}
