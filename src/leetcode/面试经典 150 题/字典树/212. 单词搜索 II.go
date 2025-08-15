package main

import "fmt"

var res []string
var existed map[string]bool
var visited [][]bool

// dfs 超时
func findWords1(board [][]byte, words []string) []string {
	res = make([]string, 0, len(words))
	existed = make(map[string]bool, 256)

	for _, word := range words {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				visited = make([][]bool, len(board))
				for i := 0; i < len(visited); i++ {
					visited[i] = make([]bool, len(board[i]))
				}
				dfs(board, word, i, j, 0)
			}
		}
	}
	return res
}

func dfs(board [][]byte, word string, i, j, wordIndex int) {
	if visited[i][j] {
		return
	}

	if word[wordIndex] != board[i][j] {
		return
	}
	visited[i][j] = true

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
	visited[i][j] = false

}

func main() {

	// {
	// 	{'a', 'b', 'c'},
	// 	{'a', 'e', 'd'},
	// 	{'a', 'f', 'g'}
	// }
	board := [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}
	words := []string{"eaafgdcba", "eaabcdgfa"}
	res := findWords(board, words)
	fmt.Println(res)
}

// https://www.bilibili.com/video/BV1rvWpeDEEn/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166

func findWords(board [][]byte, words []string) []string {

}
