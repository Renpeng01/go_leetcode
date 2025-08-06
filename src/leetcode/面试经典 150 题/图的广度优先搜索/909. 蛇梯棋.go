package main

import "fmt"

func snakesAndLadders(board [][]int) int {
	queue := make([]int, 0, 16)
	aimIndex := len(board) * len(board)
	i, j := getPoint(1, len(board))
	queue = append(queue, board[i][j])

	cnt := 0
	for len(queue) > 0 {
		cnt++
		index := queue[0]
		queue = queue[1:]
		if index == aimIndex {
			break
		}

		i, j := getPoint(index, len(board))
		if board[i][j] != -1 {
			queue = append(queue, board[i][j])
			continue
		}
		b, e := next(index, len(board))
		for k := b; k <= e; k++ {
			i, j = getPoint(k, len(board))
			queue = append(queue, board[i][j])
		}
	}
	return cnt
}

func next(index, n int) (int, int) {
	return index + 1, min(index+6, n*n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getPoint(index, n int) (int, int) {
	i := n - 1 - (index-1)/n
	j := index % n
	if j != 0 {
		j--
	}

	fmt.Println("11", j)
	if i%2 == 0 {
		j = n - 1 - j
		fmt.Println("22", n)
	}
	return i, j
}

func main() {

	var i, j int
	// i, j = getPoint(1, 6)
	// fmt.Println(i, j)

	// i, j = getPoint(8, 6)
	// fmt.Println(i, j)

	i, j = getPoint(12, 6)
	fmt.Println(i, j)
}
