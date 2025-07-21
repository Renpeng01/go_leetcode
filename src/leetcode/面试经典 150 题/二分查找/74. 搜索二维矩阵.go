package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	lo := 0
	hi := len(matrix)*len(matrix[0]) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		r := mid / len(matrix[0])
		c := mid % len(matrix[0])

		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3

	// matrix := [][]int{{1}}
	// target := 1

	res := searchMatrix(matrix, target)
	fmt.Println("res:", res)
}
