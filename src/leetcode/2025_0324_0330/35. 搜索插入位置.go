package main

import "fmt"

func searchInsert(nums []int, target int) int {

	h := 0
	t := len(nums) - 1

	mid := 0
	// f := false
	ans := len(nums)
	for h <= t {
		mid = h + (t-h)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= target {
			t = mid - 1
			ans = mid

		} else {
			h = mid + 1
		}
	}

	return ans
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
