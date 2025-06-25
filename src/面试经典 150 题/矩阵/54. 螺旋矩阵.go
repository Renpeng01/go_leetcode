package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) <= 1 {
		return matrix[0]
	}

	row := len(matrix)
	col := len(matrix[0])

	eleCnt := row * col
	res := make([]int, 0, eleCnt)

	round := 0
	cnt := 0
	for cnt < eleCnt {
		// fmt.Printf("cnt: %+v, eleCnt: %+v\n", cnt, eleCnt)
		// 上
		for i := round; i <= col-round-1; i++ {
			if cnt > eleCnt {
				break
			}
			res = append(res, matrix[round][i])
			cnt++
		}
		// 右
		for i := round + 1; i < row-round-1; i++ {
			if cnt > eleCnt {
				break
			}
			res = append(res, matrix[i][col-round-1])
			cnt++
		}
		// 下
		for i := col - round - 1; i >= round; i-- {
			if cnt > eleCnt {
				break
			}
			res = append(res, matrix[row-round-1][i])
			cnt++
		}
		// 左
		for i := row - round - 2; i >= round+1; i-- {
			if cnt > eleCnt {
				break
			}
			res = append(res, matrix[i][round])
			cnt++
		}
		round++
	}
	return res
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}
