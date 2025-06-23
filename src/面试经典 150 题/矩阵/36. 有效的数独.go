package main

// 优化点，改为一次遍历
func isValidSudoku(board [][]byte) bool {
	rows := make(map[byte]bool, 9)
	cols := make(map[byte]bool, 9)
	square := make(map[byte]bool, 9)

	for i := 0; i < 9; i++ { // row
		for j := 0; j < 9; j++ { // col

			if board[i][j] == '.' {
				continue
			}
			if rows[board[i][j]] {
				return false
			}
			rows[board[i][j]] = true
		}
		rows = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ { // col
		for j := 0; j < 9; j++ { // row
			if board[j][i] == '.' {
				continue
			}
			if cols[board[j][i]] {
				return false
			}
			cols[board[j][i]] = true
		}
		cols = make(map[byte]bool, 9)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for r := i * 3; r < i*3+3; r++ {
				for c := j * 3; c < j*3+3; c++ {
					if board[r][c] == '.' {
						continue
					}
					if square[board[r][c]] {
						return false
					}
					square[board[r][c]] = true
				}

			}
			square = make(map[byte]bool, 9)
		}
	}
	return true
}
