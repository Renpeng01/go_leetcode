package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	res := make([][]int, 0, len(intervals))
	start := -1
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] || intervals[i][0] > newInterval[1] {
			res = append(res, intervals[i])
			start = i
		}
		break
	}

	// fmt.Printf("start: %+v,res: %+v\n", start, res)
	item := []int{newInterval[0], newInterval[1]}
	otherStart := -1
	for i := start + 1; i < len(intervals); i++ {
		if item[1] >= intervals[i][0] {
			item[0] = min(item[0], intervals[i][0])
			item[1] = max(item[1], intervals[i][1])
			otherStart = i
			continue
		}
		res = append(res, item)
		break
	}
	// fmt.Printf("otherStart: %+v,res: %+v\n", otherStart, res)

	for i := otherStart + 1; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}
	// fmt.Println(res)
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

func main() {
	// intervals := [][]int{{1, 3}, {6, 9}}
	// newInterval := []int{2, 5}

	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}

	res := insert(intervals, newInterval)
	fmt.Println(res)
}
