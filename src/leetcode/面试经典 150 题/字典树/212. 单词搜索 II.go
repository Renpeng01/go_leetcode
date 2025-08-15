package main

var res []string
var existed map[string]bool
var visited map[[2]int]bool

func findWords(board [][]byte, words []string) []string {
	res = make([]string, 0, len(words))
	existed = make(map[string]bool, 256)

	for _, word := range words {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				visited = make(map[[2]int]bool, 256)
				dfs(board, word, i, j, 0)
			}
		}
	}
	return res
}

func dfs(board [][]byte, word string, i, j, wordIndex int) {
	if visited[[2]int{i, j}] {
		return
	}

	if word[wordIndex] != board[i][j] {
		return
	}

	visited[[2]int{i, j}] = true

	if wordIndex == len(word)-1 {

		if !existed[word] {
			res = append(res, word)
			existed[word] = true
		}

		return
	}

	// 上
	if i > 0 {
		dfs(board, word, i-1, j, wordIndex+1)
	}

	// 下
	if i < len(board)-1 {
		dfs(board, word, i+1, j, wordIndex+1)
	}

	// 左
	if j > 0 {
		dfs(board, word, i, j-1, wordIndex+1)
	}
	// 右
	if j < len(board[0])-1 {
		dfs(board, word, i, j+1, wordIndex+1)
	}

}
