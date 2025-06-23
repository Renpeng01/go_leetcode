package main

// 通过 但性能不佳
func setZeroes1(matrix [][]int) {

	res := make([][]int, 0, len(matrix)*len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {

				// row set 0
				for c := 0; c < len(matrix[0]); c++ {
					// matrix[i][c] = 0
					res = append(res, []int{i, c})
				}
				// rol set 0
				for r := 0; r < len(matrix); r++ {
					// matrix[r][j] = 0
					res = append(res, []int{r, j})
				}

			}
		}
	}

	for _, v := range res {
		matrix[v[0]][v[1]] = 0
	}
}

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true

			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 使用常量空间
