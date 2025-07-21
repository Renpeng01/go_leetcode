package main

import "fmt"

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

// 第一个大于等于target的index
func binarySearch3(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}

// 最后一个大于等于target
func binarySearch4(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] <= target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			l = mid + 1
		} else {
			r = mid - 1

		}
	}
	return -1
}

func main() {
	// nums := []int{1, 2, 3, 3, 3, 4}
	nums := []int{1, 1, 1, 1, 1}
	index1 := binarySearch2(nums, 2)
	fmt.Println(index1)
}
