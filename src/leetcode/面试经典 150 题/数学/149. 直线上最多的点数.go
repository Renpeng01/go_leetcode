package main

import "strconv"

func maxPoints(points [][]int) int {

	ans := 0
	for i := 0; i < len(points); i++ {
		h := make(map[string]int, 16)
		maxPointCnt := 0
		for j := i + 1; j < len(points); j++ {
			x1 := points[i][0]
			y1 := points[i][1]

			x2 := points[j][0]
			y2 := points[j][1]
			key := cal(x1, y1, x2, y2)
			h[key]++
			maxPointCnt = max(maxPointCnt, h[key])
		}
		ans = max(ans, maxPointCnt+1)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cal(x1, y1, x2, y2 int) string {

	dx := abs(x1 - x2)
	dy := abs(y1 - y2)

	k := gcd(dx, dy)

	res := strconv.Itoa(dx/k) + "_" + strconv.Itoa(dy/k)
	if (x1-x2) > 0 && (y1-y2) < 0 || (x1-x2) < 0 && (y1-y2) > 0 {
		return "-" + res
	}

	return res

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
