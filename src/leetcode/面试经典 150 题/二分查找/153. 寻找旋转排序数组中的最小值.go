package main

func findMin(nums []int) int {
	reverseIndex := -1
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
			return nums[reverseIndex]

		}
	}

	return nums[0]
}
