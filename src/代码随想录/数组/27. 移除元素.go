package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	cnt := 0

	start := -1
	end := -1

	for i, v := range nums {
		if v == val {
			cnt++
			if start == -1 {
				start = i
			}
			end = i
		}
	}

	if end == -1 {
		return len(nums)
	}

	sameLen := end - start + 1
	tail := len(nums) - 1
	i := 0
	for end <= len(nums)-1 && i < sameLen {
		nums[start], nums[tail] = nums[tail], nums[start]
		start++
		tail--
		i++
	}

	return len(nums) - cnt
}

func main() {
	nums := []int{2, 3, 3}
	res := removeElement(nums, 2)

	fmt.Println(res)
	fmt.Println(nums)
}
