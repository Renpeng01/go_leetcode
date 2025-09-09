package main

import (
	"fmt"
	"strconv"
)

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
			nums = append(nums, v2-v1)
		case "*":
			v1 := nums[len(nums)-1]
			v2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, v1*v2)
		case "/":
			v1 := nums[len(nums)-1]
			v2 := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			nums = append(nums, v2/v1)
		default:
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}
		// fmt.Println(nums)
	}
	return nums[0]
}

func main() {
	// tokens := []string{"2", "1", "+", "3", "*"}
	tokens := []string{"4", "13", "5", "/", "+"}
	res := evalRPN(tokens)
	fmt.Println(res)
}
