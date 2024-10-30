package main

import (
	"fmt"
	"sort"
)

func removeElement1(nums []int, val int) int {
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

func removeElement(nums []int, val int) int {
	cnt := 0

	head := 0
	tail := len(nums) - 1

	if len(nums) == 1 && nums[0] == val {
		return 0
	}

	for {
		for i := head; i < len(nums); i++ {
			if nums[i] == val {
				head = i
				break
			}

		}

		for i := tail; i >= 0; i-- {
			if nums[i] != val {
				tail = i
				break
			}

		}
		if tail <= head {
			break
		}

		nums[head], nums[tail] = nums[tail], nums[head]
		head++
		tail--
	}

	for _, v := range nums {
		if v == val {
			cnt++
		}
	}
	return len(nums) - cnt
}

func main() {
	nums := []int{4, 4, 4}
	res := removeElement(nums, 4)

	fmt.Println(res)
	fmt.Println(nums)
}
