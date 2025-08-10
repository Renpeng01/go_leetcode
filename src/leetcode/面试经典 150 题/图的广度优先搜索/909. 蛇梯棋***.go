package main

import "fmt"

func snakesAndLadders(board [][]int) int {
	aimIndex := len(board) * len(board)
	type pair struct{ id, step int }
	queue := []pair{{1, 0}}
	visited := make(map[int]struct{}, aimIndex)
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		b, e := next(p.id, len(board))
		for k := b; k <= e; k++ {
			nxt := k
			if nxt > aimIndex {
				break
			}
			i, j := getPoint(nxt, len(board))
			if board[i][j] > 0 {
				nxt = board[i][j]
			}
			if nxt == aimIndex { // 到达终点
				return p.step + 1
			}
			if _, ok := visited[nxt]; ok {
				continue
			}
			queue = append(queue, pair{nxt, p.step + 1})
			visited[nxt] = struct{}{}
		}
	}
	return -1
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

// func getPoint(id, n int) (r, c int) {
// 	r, c = (id-1)/n, (id-1)%n
// 	if r%2 == 1 {
// 		c = n - 1 - c
// 	}
// 	r = n - 1 - r
// 	return
// }

func main() {

	var i, j int
	i, j = getPoint(1, 6)
	fmt.Println(i, j)

	i, j = getPoint(8, 6)
	fmt.Println(i, j)

	i, j = getPoint(36, 6)
	fmt.Println(i, j)
}
