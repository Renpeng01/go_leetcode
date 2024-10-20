package main

import "math"

// 超时
func trap1(height []int) int {
	maxH := 0
	for _, v := range height {
		if v > maxH {
			maxH = v
		}
	}

	total := 0
	for i := 1; i <= maxH; i++ {
		start := -1
		target := false
		for j, v := range height {
			if v >= i {
				if target {
					total += i - start - 1
					start = j
					target = false
				} else {
					start = j
				}
			} else {
				if start != -1 {
					target = true
				}
			}
		}
	}
	return total
}

func trap2(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height)-1; i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(height[i])))
	}

	for i := len(height) - 2; i > 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(height[i])))
	}

	res := 0

	for i := 1; i < len(height)-1; i++ {
		c := int(math.Min(float64(leftMax[i]), float64(rightMax[i]))) - height[i]
		res += c
	}

	return res
}

func trap(height []int) int {
	stack := make([]int, 0, len(height))
	res := 0
	for i, v := range height {
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			l := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				h := int(math.Min(float64(v), float64(height[stack[len(stack)-1]])))
				c := (h - l) * (i - stack[len(stack)-1] - 1)
				res += c

			}

		}
		stack = append(stack, i)
	}
	return res
}
