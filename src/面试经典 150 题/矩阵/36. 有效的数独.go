package main

func isValidSudoku(board [][]byte) bool {
	rows := make(map[byte]bool, 9)
	cols := make(map[byte]bool, 9)
	square := make(map[byte]bool, 9)

	for i := 0; i < 9; i++ { // row
		for j := 0; j < 9; j++ { // col
			if rows[board[i][j]] {
				return false
			}
			rows[board[i][j]] = true
		}
		rows = make(map[byte]bool, 9)
	}

	for i := 0; i < 9; i++ { // col
		for j := 0; j < 9; j++ { // row
			if cols[board[j][i]] {
				return false
			}
			cols[board[j][i]] = true
		}
		cols = make(map[byte]bool, 9)
	}

	for m := 0; m < 9; m++ {
		step := m * 3
		for i := step; i < step+3; i++ {
			for j := step; j < step+3; j++ {
				if square[board[i][j]] {
					return false
				}
				square[board[i][j]] = true
			}
		}
		square = make(map[byte]bool, 9)
	}
	return true
}
