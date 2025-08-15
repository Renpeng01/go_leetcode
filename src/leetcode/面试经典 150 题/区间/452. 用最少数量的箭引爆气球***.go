package main

import "sort"

// https://www.bilibili.com/video/BV1SA41167xe/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
func findMinArrowShots(points [][]int) int {

	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			res++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}

	}
	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
