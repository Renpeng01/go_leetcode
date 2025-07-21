package main

import "math"

func findPeakElement(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] > get(mid-1, nums) && nums[mid] > get(mid+1, nums) {
			return mid
		}

		if get(mid-1, nums) < nums[mid] && nums[mid] < get(mid+1, nums) {
			l = mid + 1
			continue

		}

		if get(mid-1, nums) > nums[mid] && nums[mid] > get(mid+1, nums) {
			r = mid - 1
			continue
		}

		if get(mid-1, nums) > nums[mid] && nums[mid] < get(mid+1, nums) {
			r = mid - 1
			continue
		}

	}
	return -1
}

func get(i int, nums []int) int {
	if i == -1 || i == len(nums) {
		return math.MinInt
	}

	return nums[i]

}
