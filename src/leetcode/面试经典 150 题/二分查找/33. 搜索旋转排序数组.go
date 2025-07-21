package main

import "fmt"

func search(nums []int, target int) int {

	reverseIndex := -1

	start := 0
	end := len(nums) - 1
	if nums[0] > nums[len(nums)-1] {
		l := 0
		r := len(nums) - 1
		for l <= r {
			mid := l + (r-l)/2

			if mid+1 <= len(nums)-1 && mid-1 >= 0 && nums[mid-1] > nums[mid] && nums[mid] < nums[mid+1] {
				reverseIndex = mid
				break
			}

			// 在左侧
			if nums[0] <= nums[mid] {
				if mid+1 <= len(nums)-1 && nums[mid] > nums[mid+1] {
					reverseIndex = mid + 1
					break
				}
				l = mid + 1
				continue
			}

			// 在右侧
			if nums[mid] <= nums[len(nums)-1] {
				if mid-1 >= 0 && nums[mid-1] > nums[mid] {
					reverseIndex = mid - 1
					break
				}
				r = mid - 1
				continue
			}
		}

		if reverseIndex != -1 {

			fmt.Println("reverseIndex: ", reverseIndex)

			if target >= nums[0] && target <= nums[reverseIndex-1] {
				start = 0
				end = reverseIndex - 1
			} else {
				start = reverseIndex
				end = len(nums) - 1

			}
		}
	}

	// fmt.Println(start, end, target)
	if target < nums[start] || target > nums[end] {
		// fmt.Println(111111)
		return -1
	}

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func main() {

	// nums := []int{4, 5, 6, 7, 0, 1, 2}

	nums := []int{3, 1}
	index := search(nums, 0)
	fmt.Println(index)

}
