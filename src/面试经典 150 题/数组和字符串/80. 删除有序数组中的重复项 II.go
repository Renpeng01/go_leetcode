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
