package main

import (
	"fmt"
)

func insert1(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	res := make([][]int, 0, len(intervals))
	start := -1
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
			start = i
		} else if intervals[i][0] > newInterval[1] {
			// res = append(res, newInterval)
			// res = append(res, intervals[i])
			// start = i
			// break
			res = append(res, newInterval)
			for j := i; j < len(intervals); j++ {
				res = append(res, intervals[j])
			}
			return res
		}
		break
	}

	fmt.Printf("start: %+v,res: %+v\n", start, res)
	item := []int{newInterval[0], newInterval[1]}
	fmt.Printf("item: %+v\n", item)
	otherStart := start
	for i := start + 1; i < len(intervals); i++ {
		if item[1] < intervals[i][0] {
			otherStart = i - 1
			break
		}
		if item[1] >= intervals[i][0] {
			item[0] = min(item[0], intervals[i][0])
			item[1] = max(item[1], intervals[i][1])
			otherStart = i
			continue
		}
		break
	}

	res = append(res, item)
	fmt.Printf("otherStart: %+v,res: %+v\n", otherStart, res)

	for i := otherStart + 1; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}
	fmt.Println(res)
	return res
}

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
