package main

func solve(board [][]byte) {
	// 上
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' {
			dfs(board, 0, i)
		}
	}

	// 下
	for i := 0; i < len(board[0]); i++ {
		if board[len(board)-1][i] == 'O' {
			dfs(board, len(board)-1, i)
		}
	}

	// 左
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
	}

	// 右
	for i := 0; i < len(board); i++ {
		if board[i][len(board[0])-1] == 'O' {
			dfs(board, i, len(board[0])-1)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'i' {
				board[i][j] = 'O'
			}
		}
	}

}

func dfs(board [][]byte, i, j int) {
	board[i][j] = 'i'
	// 上
	if i-1 >= 0 && board[i-1][j] == 'O' {
		dfs(board, i-1, j)
	}
	// 下
	if i+1 < len(board) && board[i+1][j] == 'O' {
		dfs(board, i+1, j)
	}
	// 左
	if j-1 >= 0 && board[i][j-1] == 'O' {
		dfs(board, i, j-1)
	}
	// 右
	if j+1 < len(board[0]) && board[i][j+1] == 'O' {
		dfs(board, i, j+1)
	}
}
