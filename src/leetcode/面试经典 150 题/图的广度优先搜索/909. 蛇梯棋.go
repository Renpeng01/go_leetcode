package main

import "fmt"

func snakesAndLadders(board [][]int) int {
	queue := make([]int, 0, 16)
	aimIndex := len(board) * len(board)
	queue = append(queue, 1)

	visited := make(map[int]struct{}, 256)

	cnt := 0
	for len(queue) > 0 {
		index := queue[0]
		queue = queue[1:]

		if _, ok := visited[index]; ok {
			continue
		}
		visited[index] = struct{}{}
		if index == aimIndex {
			break
		}
		i, j := getPoint(index, len(board))
		if board[i][j] != -1 {
			queue = append(queue, board[i][j])
			cnt++
			continue
		}
		b, e := next(index, len(board))
		for k := b; k <= e; k++ {
			queue = append(queue, k)
		}
		cnt++
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
	j := (index - 1) % n
	if i%2 == 0 {
		j = n - 1 - j
	}
	return i, j
}

func main() {

	var i, j int
	i, j = getPoint(1, 6)
	fmt.Println(i, j)

	i, j = getPoint(8, 6)
	fmt.Println(i, j)

	i, j = getPoint(36, 6)
	fmt.Println(i, j)
}
