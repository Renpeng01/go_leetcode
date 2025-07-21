package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	lo := 0
	hi := len(matrix) - 1

	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if matrix[mid][hi] == target {
			return true
		}
		if matrix[mid][0] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}

	}

	// if hi < 0 {
	// 	return false
	// }

	// fmt.Println("hi: ", hi)

	if matrix[mid][hi] == target {
		return true
	}

	row := hi

	lo = 0
	hi = len(matrix[0]) - 1

	for lo <= hi {
		mid = lo + (hi-lo)/2
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

func main() {
	// matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	// target := 11

	matrix := [][]int{{1}}
	target := 1

	res := searchMatrix(matrix, target)
	fmt.Println("res:", res)
}
