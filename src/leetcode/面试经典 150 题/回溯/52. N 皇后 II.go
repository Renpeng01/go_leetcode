package main

import "fmt"

var res [][][]int
var path [][]int

func totalNQueens(n int) int {
	res = make([][][]int, 0, 16)
	path = make([][]int, 0, 16)
	backtracking(0, n)
	fmt.Println(res)
	return len(res)
}

func backtracking(row, n int) {
	if row == n {
		tmp := make([][]int, 0, 16)
		for _, pair := range path {
			tmp = append(tmp, pair)
		}
		res = append(res, tmp)
		return
	}

	for i := 1; i <= n; i++ {
		x := row
		y := i
		valid := true
		for _, pair := range path {
			x1 := pair[0]
			y1 := pair[1]
			if x == x1 || y == y1 || (x-y) == (x1-y1) || (x+y) == (x1+y1) {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		path = append(path, []int{row, i})
		backtracking(row+1, n)
		path = path[:len(path)-1]
	}
}

func main() {
	res := totalNQueens(4)
	fmt.Println(res)
}
