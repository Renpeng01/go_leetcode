package main

import "fmt"

var path [][2]int
var pathSet map[[2]int]struct{}

var steps [][]int = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	path = make([][2]int, 0, 16)
	pathSet = make(map[[2]int]struct{}, 16)

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
		path = append(path, [2]int{i, j})
		pathSet[[2]int{i, j}] = struct{}{}
		if len(path) == len(word) {
			return true
		}

		for _, step := range steps {
			if _, ok := pathSet[[2]int{i + step[0], j + step[1]}]; !ok && i+step[0] >= 0 && i+step[0] < len(board) && j+step[1] >= 0 && j+step[1] < len(board[0]) {
				if backtracking(board, i+step[0], j+step[1], n+1, word) {
					return true
				}
			}
		}
		delete(pathSet, path[len(path)-1])
		path = path[:len(path)-1]

	}
	return false
}

func main() {

	board := [][]byte{{'a', 'b'}}
	word := "ba"

	res := exist(board, word)
	fmt.Println("res: ", res)
}
