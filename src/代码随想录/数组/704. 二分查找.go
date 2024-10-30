package main

func search(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1

	for head <= tail {
		mid := head + (tail-head)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			tail = mid - 1
		} else {
			head = mid + 1
		}

	}
	return -1
}
