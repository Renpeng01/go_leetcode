package main

import (
	"fmt"
)

var res [][]int
var path []int
var curSum int

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0, 16)
	path = make([]int, 0, 16)
	curSum = 0
	backtracking(candidates, target, 0)

	return res
}

func backtracking(candidates []int, target, startIndex int) {
	if curSum > target {
		return
	}
	if curSum == target {
		tmp := make([]int, 0, len(path))
		for _, v := range path {
			tmp = append(tmp, v)
		}
		res = append(res, tmp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		path = append(path, candidates[i])
		curSum += candidates[i]
		backtracking(candidates, target, i)
		curSum -= path[len(path)-1]
		path = path[:len(path)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
