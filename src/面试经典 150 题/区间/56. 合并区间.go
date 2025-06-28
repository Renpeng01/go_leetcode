package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// fmt.Println(intervals)

	res := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); {

		item := make([]int, 2)
		item[0] = intervals[i][0]
		item[1] = intervals[i][1]

		e := i
		for j := i + 1; j < len(intervals); j++ {
			if item[1] < intervals[j][0] {
				break
			}
			item[1] = max(intervals[j][1], item[1])
			e = j
		}

		res = append(res, item)
		i = e + 1
	}
	return res

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	intervals := [][]int{{1, 4}, {0, 2}, {3, 5}}
	res := merge(intervals)
	fmt.Println(res)
}
