package main

import "fmt"

var res [][][2]int
var path [][2]int

var pathSet map[[2]int]struct{}

func exist(board [][]byte, word string) bool {
	res = make([][][2]int, 0, 16)
	path = make([][2]int, 0, 16)
	pathSet = make(map[[2]int]struct{}, 16)
	backtracking(board, 0, 0, 0, word)
	fmt.Println(res)
	return len(res) > 0
}

func backtracking(board [][]byte, ii, jj, n int, word string) {
	for i := ii; i < len(board); i++ {
		for j := jj; j < len(board[0]); j++ {
			if _, ok := pathSet[[2]int{i, j}]; ok {
				break
			}

			if board[i][j] != word[n] {
				break
			}

			path = append(path, [2]int{i, j})
			pathSet[[2]int{i, j}] = struct{}{}

			if n+1 == len(word) {
				tmp := make([][2]int, 0, 16)
				for _, pair := range path {
					tmp = append(tmp, pair)
				}
				res = append(res, tmp)
				return
			}
			// 上
			if i-1 >= 0 {
				if _, ok := pathSet[[2]int{i - 1, j}]; !ok {
					backtracking(board, i-1, j, n+1, word)
				}
			}

			// 下
			if i+1 < len(board) {
				if _, ok := pathSet[[2]int{i + 1, j}]; !ok {
					backtracking(board, i+1, j, n+1, word)
				}
			}
			// 左
			if j-1 >= 0 {
				if _, ok := pathSet[[2]int{i, j - 1}]; !ok {
					backtracking(board, i, j-1, n+1, word)
				}
			}
			// 右
			if j+1 < len(board[0]) {
				if _, ok := pathSet[[2]int{i, j + 1}]; !ok {
					backtracking(board, i, j+1, n+1, word)
				}
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
