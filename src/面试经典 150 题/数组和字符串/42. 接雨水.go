package main

import (
	"fmt"
)

// 超时
func trap(height []int) int {
	base := 1
	total := 0

	for {
		b := -1
		heighterCnt := 0
		for i := 0; i < len(height); i++ {
			if height[i] >= base {
				if b < 0 {
					b = i
				} else {
					total += (i - b) - 1
					b = i
				}
				heighterCnt++
			}
		}
		if heighterCnt <= 1 {
			break
		}
		base++
	}
	return total
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}

func trap(height []int) int {

	stack := make([]int, 0, len(height))
	res := 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) <= 0 {
				break
			}

			minH := min(height[stack[len(stack)-1]], h) - height[top]
			with := i - 1 - stack[len(stack)-1]
			res += minH * with

		}
		stack = append(stack, i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
