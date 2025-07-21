package main

func searchRange(nums []int, target int) []int {
	v1 := binarySearch1(nums, target)
	v2 := binarySearch2(nums, target)
	return []int{v1, v2}

}

// 第一个等于target的index
func binarySearch1(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {

			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// 最后一个等于target的index
func binarySearch2(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
