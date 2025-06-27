package main

func twoSum(nums []int, target int) []int {

	s := make(map[int]int, len(nums))

	res := make([]int, 2)
	for i, v := range nums {
		if j, ok := s[target-v]; ok {
			res[0], res[1] = i, j
			return res
		}
		s[v] = i
	}
	return res

}
