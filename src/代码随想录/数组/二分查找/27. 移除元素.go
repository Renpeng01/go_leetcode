package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println("sorted: ", nums)

	// 找到第一个等于val的位置
	// 如果没有找到 return 0
	mid, left, right := -1, 0, len(nums)-1
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == val {
			if mid > 0 && nums[mid-1] == val {
				right = mid - 1
			}
			return mid
		} else if nums[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("mid: ", mid)

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
	fmt.Println("cnt: ", cnt)

	n := min(cnt, len(nums)-1-cnt)
	for i := 0; i <= n; i++ {
		nums[mid+i], nums[mid+i+cnt] = nums[mid+i+cnt], nums[mid+i]
	}
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println("res: ", res)
}
