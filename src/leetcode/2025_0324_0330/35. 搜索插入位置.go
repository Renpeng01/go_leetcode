package main

import "fmt"

func searchInsert(nums []int, target int) int {

	if target <= nums[0] {
		return 0
	}

	h := 0
	t := len(nums) - 1

	mid := 0
	for h < t {
		mid = h + (t-h)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			t = mid - 1
		} else {
			h = mid + 1
		}
	}
	return mid
}

func main() {
	nums := [][]int{{1, 3, 5, 6}, {1, 3, 5, 6}, {1, 3, 5, 6}}
	targets := []int{5, 2, 7} // 2 1 4

	// nums := [][]int{{1, 3}} // 1
	// targets := []int{2}

	for i, vals := range nums {
		res := searchInsert(vals, targets[i])
		fmt.Println(res)
	}
}
