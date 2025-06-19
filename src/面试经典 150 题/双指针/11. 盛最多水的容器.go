package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := min(height[left], height[right]) * (right - left)

	for left < right {
		nextLeft := left + 1
		nextRigth := right - 1
		if nextLeft < right && height[left] < height[right] && height[nextLeft] > height[left] {
			if min(height[nextLeft], height[right])*(right-nextLeft) > maxArea {
				maxArea = min(height[nextLeft], height[right]) * (right - nextLeft)
			}
			left++
			continue
		}

		if nextRigth > left && height[right] < height[left] && height[nextRigth] > height[right] {
			if min(height[nextRigth], height[left])*(nextRigth-left) > maxArea {
				maxArea = min(height[nextRigth], height[left]) * (nextRigth - left)

			}
			right--
			continue
		}
		left++
		right--
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
