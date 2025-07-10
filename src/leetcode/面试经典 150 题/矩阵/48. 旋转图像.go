package main

import "fmt"

func rotate(matrix [][]int) {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i], matrix[j][len(matrix)-1-i] = matrix[j][len(matrix)-1-i], matrix[j][i]
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	PrintMatrix(matrix)

}

func PrintMatrix(matrix [][]int) {
	for _, vals := range matrix {
		for _, v := range vals {
			fmt.Printf("%+v ", v)
		}
		fmt.Println()
	}
}
