package main

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)

	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == target {
			res[0] = left + 1
			res[1] = right + 1
			return res
		}

		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return res
}
