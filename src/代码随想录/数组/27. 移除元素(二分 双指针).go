package main

import (
	"fmt"
	"sort"
)

// 二分方法
func removeElement1(nums []int, val int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 找到第一个等于val的位置
	// 如果没有找到 return 0
	mid, left, right := -1, 0, len(nums)-1
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == val {
			if mid > 0 && nums[mid-1] == val {
				right = mid - 1
				continue
			}
			break
		} else if nums[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if mid == -1 {
		return 0
	}

	cnt := 0
	for i := mid; i < len(nums); i++ {
		if nums[i] != val {
			break
		}
		cnt++
	}

	for i := 0; i <= len(nums)-1-mid-cnt; i++ {
		nums[mid+i] = nums[mid+i+cnt]
	}
	return len(nums) - cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums := []int{2, 2, 2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println("res: ", res)
}

// 双指针
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
