package main

import "math"

// 性能差
func maxArea1(height []int) int {
	max := 0
	leftMax := 0
	for i := 0; i < len(height); i++ {
		if height[i] > leftMax {
			leftMax = height[i]
		} else {
			continue
		}
		rightMax := 0
		for j := len(height) - 1; j > i; j-- {
			if height[j] > rightMax {
				rightMax = height[j]
			} else {
				continue
			}

			h := math.Min(float64(height[j]), float64(height[i]))
			if int(h)*(j-i) > max {
				max = int(h) * (j - i)
			}
		}
	}

	return max
}

func maxArea(height []int) int {
	head := 0
	tail := len(height) - 1
	max := 0

	for head < tail {
		h := math.Min(float64(height[head]), float64(height[tail]))
		if int(h)*(tail-head) > max {
			max = int(h) * (tail - head)
		}

		if height[head] > height[tail] {
			tail--
		} else {
			head++
		}
	}

	return max
}
