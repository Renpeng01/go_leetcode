package main

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
	return 0, 0
}

func getIndex(i, j int) int {
	return 0
}
