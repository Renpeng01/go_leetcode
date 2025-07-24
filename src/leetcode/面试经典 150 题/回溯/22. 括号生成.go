package main

var res []string
var path []byte

func generateParenthesis(n int) []string {
	res = make([]string, 0, 16)
	path = make([]byte, 0, 16)
	backtracking(0, n)
	return res
}

func backtracking(cur, n int) {
	if cur == n {
		return
	}

}
