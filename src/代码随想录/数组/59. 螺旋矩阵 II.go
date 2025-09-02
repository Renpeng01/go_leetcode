package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	val := 1
	round := (n + 1) / 2
	for i := 0; i < round; i++ {
		// 上
		for j := i; j < n-i; j++ {
			res[i][j] = val
			val++
		}
		// 右
		for j := i + 1; j < n-i-1; j++ {
			res[j][n-i-1] = val
			val++
		}

		if res[n-i-1][n-i-1] != 0 {
			break
		}

		// 下
		for j := i; j < n-i; j++ {
			res[n-i-1][n-j-1] = val
			val++
		}

		// 左
		for j := i + 1; j < n-i-1; j++ {
			res[n-j-1][i] = val
			val++
		}
	}
	return res
}

func main() {
	res := generateMatrix(4)
	fmt.Println(res)
}
