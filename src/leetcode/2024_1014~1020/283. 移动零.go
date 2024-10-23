package main

import "fmt"

func moveZeroes(nums []int) {
	targetIndex := -1
	if len(nums) > 0 {
		if nums[0] == 0 {
			targetIndex = 0
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] != 0 {
			if targetIndex >= 0 {
				nums[targetIndex], nums[i] = nums[i], nums[targetIndex]
				targetIndex++
			}

		} else {
			if targetIndex == -1 {
				targetIndex = i
			}
		}
	}
}

func main() {
	// data := []int{0, 1, 0, 3, 12}
	data := []int{0}
	moveZeroes(data)
	fmt.Println(data)
}
