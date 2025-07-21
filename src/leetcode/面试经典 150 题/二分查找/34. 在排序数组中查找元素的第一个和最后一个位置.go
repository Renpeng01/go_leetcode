package main

func searchRange(nums []int, target int) []int {

}

func binarySearch1(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l

}

// 最后一个等于target的index
func binarySearch2(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
