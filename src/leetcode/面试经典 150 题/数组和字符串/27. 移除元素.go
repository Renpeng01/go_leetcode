package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {

	tail := findTail(nums, 0, len(nums)-1, val)
	head := 0
	for head <= tail {
		if nums[head] == val {
			nums[head], nums[tail] = nums[tail], nums[head]
			tail = findTail(nums, head, tail-1, val)
			head++
			continue
		}
		head++
	}

	return head

}

func findTail(nums []int, head, tail, val int) int {
	res := -1
	for i := tail; i >= head; i-- {
		if nums[i] != val {
			res = i
			break
		}
	}

	return res
}

func removeElement1(nums []int, val int) int {
	tail := len(nums) - 1
	head := 0
	for head < tail {

		if nums[head] == val {
			nums[head] = nums[tail]
			tail--
		} else {
			head++
		}
	}

	return head

}

func main() {
	nums := []int{5, 5}
	res := removeElement(nums, 5)
	fmt.Println(res, nums)
}
