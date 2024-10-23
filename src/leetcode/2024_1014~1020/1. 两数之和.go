package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		m[v] = i
	}

	res := make([]int, 2)

	for i, v := range nums {
		index, ok := m[target-v]
		if ok && i != index {
			res[0] = i
			res[1] = index
			return res
		}
	}

	return nil
}
