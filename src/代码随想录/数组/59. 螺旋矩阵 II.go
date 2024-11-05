package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	curNum := 1
	for i := 0; i < n/2; i++ {
		// 上
		for j := i; j < n-i; j++ {
			res[i][j] = curNum
			curNum++
		}
		// 右
		for m := i + 1; m < n-i-1; m++ {
			res[m][n-1-i] = curNum
			curNum++
		}

		// 下
		for k := n - 1 - i; k >= i; k-- {
			res[n-1-i][k] = curNum
			curNum++
		}
		// 左
		for q := n - 1 - i - 1; q >= i+1; q-- {
			res[q][i] = curNum
			curNum++
		}

	}

	if n%2 == 1 {
		res[n/2][n/2] = n * n
	}
	return res
}

func main() {
	n := 5
	res := generateMatrix(n)
	fmt.Println(res)
}
