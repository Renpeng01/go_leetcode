package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	lo := 0
	hi := len(matrix) - 1

	size := len(matrix)

	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			if (mid-1 >= 0 && matrix[mid-1][0] < target) || mid == 0 {
				// fmt.Println(222)
				break
			}
			hi = mid - 1
		} else {
			if (mid+1 <= size-1 && matrix[mid+1][0] > target) || mid == size-1 {
				// fmt.Println(1111)
				break
			}
			lo = mid + 1
		}
	}

	fmt.Println(mid)
	row := mid

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
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	target := 30

	// matrix := [][]int{{1}}
	// target := 1

	res := searchMatrix(matrix, target)
	fmt.Println("res:", res)
}
