package main

import (
	"fmt"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	isInsert := false
	res := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); {
		if intervals[i][0] > newInterval[1] {
			if !isInsert {
				res = append(res, newInterval)
			}
			res = append(res, intervals[i])
			isInsert = true
			i++
		} else if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
			if i == len(intervals)-1 {
				res = append(res, newInterval)
			}
			i++
		} else {
			cur := i
			for j := cur; j < len(intervals); j++ {
				if newInterval[1] >= intervals[j][0] {
					newInterval[0] = min(newInterval[0], intervals[j][0])
					newInterval[1] = max(newInterval[1], intervals[j][1])
					isInsert = true
					cur = j
					continue
				}
				break
			}
			res = append(res, newInterval)
			for m := cur + 1; m < len(intervals); m++ {
				res = append(res, intervals[m])
				cur = m
			}
			i = cur + 1
		}
	}
	return res
}

func main() {

	// intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	// newInterval := []int{4, 8}

	// intervals := [][]int{{1, 5}}
	// newInterval := []int{2, 3}

	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}

	// intervals := [][]int{{1, 5}}
	// newInterval := []int{0, 0}

	res := insert(intervals, newInterval)
	fmt.Println(res)
}
