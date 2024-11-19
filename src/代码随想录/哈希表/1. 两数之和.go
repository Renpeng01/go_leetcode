package main

func twoSum(nums []int, target int) []int {
	posSet := make(map[int][]int, len(nums))

	for i, v := range nums {
		if _, ok := posSet[v]; !ok {
			posSet[v] = make([]int, 0, 2)

		}
		posSet[v] = append(posSet[v], i)
	}

	res := make([]int, 2)
	for i, v := range nums {
		if vals, ok := posSet[target-v]; ok {
			for _, index := range vals {
				if i != index {
					res[0] = i
					res[1] = index
					return res
				}
			}

		}
	}
	return res
}
