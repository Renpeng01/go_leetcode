package main

import (
	"fmt"
	"sort"
	"strconv"
)

var res [][]int
var path []int
var curSum int
var exist map[string]struct{}

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0, 16)
	path = make([]int, 0, 16)
	exist = make(map[string]struct{}, 16)
	curSum = 0
	backtracking(candidates, target)

	return res
}

func backtracking(candidates []int, target int) {
	if curSum > target {
		return
	}
	if curSum == target {
		tmp := make([]int, 0, len(path))
		for _, v := range path {
			tmp = append(tmp, v)
		}
		sort.Ints(tmp)
		key := ""
		for _, v := range tmp {
			key = key + strconv.Itoa(v) + "_"
		}

		if _, ok := exist[key]; ok {
			return
		}

		exist[key] = struct{}{}
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(candidates); i++ {
		path = append(path, candidates[i])
		curSum += candidates[i]
		backtracking(candidates, target)
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
