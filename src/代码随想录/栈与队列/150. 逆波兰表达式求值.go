package main

import "strconv"

func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))
	for _, v := range tokens {

		switch v {
		case "+":
			v1 := nums[len(nums)-1]
			v2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, v1+v2)
		case "-":
			v1 := nums[len(nums)-1]
			v2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, v1-v2)
		case "*":
			v1 := nums[len(nums)-1]
			v2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, v1*v2)
		case "/":
			v1 := nums[len(nums)-1]
			v2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, v1/v2)
		default:
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}
	}
	return nums[0]
}
