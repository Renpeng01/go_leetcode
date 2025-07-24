package main

import "fmt"

var pathSet [][]bool

var steps [][]int = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	pathSet = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		pathSet[i] = make([]bool, len(board[i]))
	}

	fmt.Println(pathSet)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtracking(board, i, j, 0, word) {
				return true
			}
		}

	}
	return false
}

func backtracking(board [][]byte, i, j, n int, word string) bool {

	if board[i][j] == word[n] {
		pathSet[i][j] = true
		if n+1 == len(word) {
			return true
		}

		for _, step := range steps {

			if i+step[0] >= 0 && i+step[0] < len(board) && j+step[1] >= 0 && j+step[1] < len(board[0]) && !pathSet[i+step[0]][j+step[1]] {
				if backtracking(board, i+step[0], j+step[1], n+1, word) {
					return true
				}
			}
		}
		pathSet[i][j] = false

	}
	return false
}

func main() {

	board := [][]byte{{'a', 'b'}}
	word := "ba"

	res := exist(board, word)
	fmt.Println("res: ", res)
}
