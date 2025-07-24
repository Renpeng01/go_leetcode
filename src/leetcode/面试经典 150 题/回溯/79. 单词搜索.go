package main

var res [][][]int
var path [][]int

var pathSet map[[2]int]struct{}

func exist(board [][]byte, word string) bool {
	res = make([][][]int, 0, 16)
	path = make([][]int, 0, 16)
	pathSet = make(map[[2]int]struct{}, 16)
	backtracking(board, 0, 0, 0, word)
	return len(res) > 0
}

func backtracking(board [][]byte, ii, jj, n int, word string) {
	if n == len(word) {
		tmp := make([][]int, 0, 16)
		for _, pair := range path {
			tmp = append(tmp, pair)
		}
		res = append(res, tmp)
		return
	}
	for i := ii; i < len(board); i++ {
		for j := jj; j < len(board[0]); j++ {
			if _, ok := pathSet[[2]int{i, j}]; ok {
				continue
			}

			if board[i][j] != word[n] {
				continue
			}

			path = append(path, []int{i, j})
			pathSet[[2]int{i, j}] = struct{}{}
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
			path = path[:len(path)-1]
			delete(pathSet, [2]int{i, j})
		}
	}
}
