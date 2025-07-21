package main

func searchMatrix(matrix [][]int, target int) bool {
	lo := 0
	hi := len(matrix) - 1

	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}

	}

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
