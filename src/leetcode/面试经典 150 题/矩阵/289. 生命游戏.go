package main

func gameOfLife1(board [][]int) {

	offsets := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	newBard := make([][]int, len(board))
	for i := 0; i < len(newBard); i++ {
		newBard[i] = make([]int, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			c0 := 0
			c1 := 0
			for _, v := range offsets {
				if i+v[0] >= 0 && i+v[0] < len(board) && j+v[1] >= 0 && j+v[1] < len(board[0]) {
					if board[i+v[0]][j+v[1]] == 0 {
						c0++
					} else {
						c1++
					}
				}
			}

			v := board[i][j]
			if board[i][j] == 1 {
				if c1 < 2 {
					v = 0
				} else if c1 >= 2 && c1 <= 3 {
					v = 1
				} else {
					v = 0
				}
			} else {
				if c1 == 3 {
					v = 1
				}
			}
			newBard[i][j] = v
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = newBard[i][j]
		}
	}
}

func gameOfLife(board [][]int) {

	offsets := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// 0->1 2
	// 1->0 3
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			c0 := 0
			c1 := 0
			for _, v := range offsets {
				if i+v[0] >= 0 && i+v[0] < len(board) && j+v[1] >= 0 && j+v[1] < len(board[0]) {
					if board[i+v[0]][j+v[1]] == 0 || board[i+v[0]][j+v[1]] == 2 {
						c0++
					} else {
						c1++
					}
				}
			}

			v := board[i][j]
			if board[i][j] == 1 {
				if c1 < 2 {
					v = 3
				} else if c1 > 3 {
					v = 3
				}
			} else {
				if c1 == 3 {
					v = 2
				}
			}
			board[i][j] = v
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}
func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
}
