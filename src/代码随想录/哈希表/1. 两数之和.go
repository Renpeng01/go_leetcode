package main

func twoSum1(nums []int, target int) []int {
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

// 优化，不需要提前把全部的位置都存在set中，可以在遍历的同时查找结果，也能同时避免使用同一个位置元素的问题
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	posSet := make(map[int]int, len(nums))
	for i, v := range nums {
		index, ok := posSet[target-v]
		if !ok {
			posSet[v] = i
			continue
		}
		res[0] = i
		res[1] = index
		return res
	}
	return res
}
