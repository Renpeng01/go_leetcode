package main

var res []string
var existed map[string]bool

// var visited [][]bool

// dfs 超时
func findWords(board [][]byte, words []string) []string {
	res = make([]string, 0, len(words))
	existed = make(map[string]bool, 256)

	for _, word := range words {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				dfs(board, word, i, j, 0)
			}
		}
	}
	return res
}

func dfs(board [][]byte, word string, i, j, wordIndex int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' || wordIndex >= len(word) {
		return
	}

	if word[wordIndex] != board[i][j] {
		return
	}

	c := board[i][j]
	board[i][j] = '*'

	if wordIndex == len(word)-1 {
		if !existed[word] {
			res = append(res, word)
			existed[word] = true
		}
	}

	dfs(board, word, i-1, j, wordIndex+1)
	dfs(board, word, i+1, j, wordIndex+1)
	dfs(board, word, i, j-1, wordIndex+1)
	dfs(board, word, i, j+1, wordIndex+1)
	board[i][j] = c
}

// func main() {

// 	// {
// 	// 	{'a', 'b', 'c'},
// 	// 	{'a', 'e', 'd'},
// 	// 	{'a', 'f', 'g'}
// 	// }
// 	board := [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}
// 	words := []string{"eaafgdcba", "eaabcdgfa"}
// 	res := findWords(board, words)
// 	fmt.Println(res)
// }

// type treeNode struct {
// 	data map[byte]*treeNode
// 	word string
// }

// func NewtreeNode() *treeNode {
// 	return &treeNode{
// 		data: make(map[byte]*treeNode, 256),
// 	}
// }
// func (treeNode *treeNode) insert(word string) {
// 	node := treeNode
// 	for i := 0; i < len(word); i++ {
// 		if _, ok := node.data[word[i]]; !ok {
// 			node.data[word[i]] = NewtreeNode()
// 		}
// 		node = node.data[word[i]]
// 	}
// 	node.word = word
// }

// var resMap map[string]struct{}

// // https://www.bilibili.com/video/BV1rvWpeDEEn/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
// func findWords(board [][]byte, words []string) []string {
// 	resMap = make(map[string]struct{}, 256)
// 	node := NewtreeNode()
// 	for _, w := range words {
// 		node.insert(w)
// 	}

// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			dfs(board, i, j, node)
// 		}
// 	}

// 	resList := make([]string, 0, len(resMap))
// 	for k := range resMap {
// 		resList = append(resList, k)
// 	}
// 	return resList
// }

// func dfs(board [][]byte, i, j int, node *treeNode) {

// if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' || node.data[board[i][j]] == nil {
// 	return
// }

// 	newNode := node.data[board[i][j]]
// 	if len(newNode.word) > 0 {
// 		resMap[newNode.word] = struct{}{}

// 	}
// 	c := board[i][j]
// 	board[i][j] = '*'

// 	dfs(board, i-1, j, newNode)
// 	dfs(board, i+1, j, newNode)
// 	dfs(board, i, j-1, newNode)
// 	dfs(board, i, j+1, newNode)

// 	board[i][j] = c
// }
