package main

func searchMatrix(matrix [][]int, target int) bool {

	r := 0
	c := 0
	for r < len(matrix) && r > 0 && c < len(matrix[0]) && c > 0 {
		rLo := 0
		rHi := len(matrix)
		rMid := rLo + (rHi-rLo)/2
		for rLo <= rHi {
			if matrix[0][rMid] == target {
				return true
			}
			if target > matrix[0][rMid] {
				rLo = rMid + 1
			} else {
				rHi = rMid - 1
			}
		}
		r = rLo

		cLo := 0
		cHi := len(matrix[0])
		cMid := cLo + (cHi-rLo)/2
		for cLo <= cHi {
			if matrix[cMid][0] == target {
				return true
			}
			if target > matrix[0][cMid] {
				cLo = cMid + 1
			} else {
				cHi = cMid - 1
			}
		}
		c = cLo
	}

	return false
}

func main() {

}
