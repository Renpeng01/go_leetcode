package main

import "fmt"

var res [][][2]int
var path [][2]int
var pathSet map[[2]int]struct{}

var steps [][]int = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	res = make([][][2]int, 0, 16)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			path = make([][2]int, 0, 16)
			pathSet = make(map[[2]int]struct{}, 16)
			backtracking(board, i, j, 0, word)
		}
	}
	fmt.Println(res)
	return len(res) > 0
}

func backtracking(board [][]byte, i, j, n int, word string) {
	for _, step := range steps {
		if board[i][j] == word[n] {
			path = append(path, [2]int{i, j})
			pathSet[[2]int{i, j}] = struct{}{}

			if len(path) == len(word) {
				fmt.Println(path)
				tmp := make([][2]int, 0, 16)
				for _, pair := range path {
					tmp = append(tmp, pair)
				}
				res = append(res, tmp)
				return
			}

			if _, ok := pathSet[[2]int{i + step[0], j + step[1]}]; !ok && i+step[0] >= 0 && i+step[0] < len(board) && j+step[1] > 0 && j+step[1] < len(board[0]) {
				backtracking(board, i+step[0], j+step[1], n+1, word)
			}
			delete(pathSet, path[len(path)-1])
			path = path[:len(path)-1]
		}
	}
}

func main() {
	board := [][]byte{{'b'}, {'a'}, {'b'}, {'b'}, {'a'}}
	word := "baa"

	res := exist(board, word)
	fmt.Println("res: ", res)
}
