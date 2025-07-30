package main

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	dp[0][0] = 0
	if matrix[0][0] == '1' {
		dp[0][0] = 1
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				// dp[i][j] = dp[i-1][j]

				if 
			} else {

			}
		}
	}

	return dp[len(matrix)-1][len(matrix[0])-1]
}

func max(a,b int) int{
	if a > b{
		return a
	}
	return b
}