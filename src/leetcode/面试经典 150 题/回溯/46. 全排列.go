package main

var res [][]int
var path []int
var block map[int]struct{}

func permute(nums []int) [][]int {

	res = make([][]int, 0, 16)
	path = make([]int, 0, 16)

	block = make(map[int]struct{}, 16)
	backtracking(nums)
	return res

}

func backtracking(nums []int) {
	if len(path) == len(nums) {
		tmp := make([]int, 0, len(path))
		for _, v := range path {
			tmp = append(tmp, v)
		}
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := block[nums[i]]; ok {
			continue
		}
		path = append(path, nums[i])
		block[nums[i]] = struct{}{}
		backtracking(nums)
		delete(block, nums[i])
		path = path[:len(path)-1]
	}
}
