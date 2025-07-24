package main

var res []string
var path []byte
var done map[string]struct{}

func generateParenthesis(n int) []string {
	res = make([]string, 0, 16)
	path = make([]byte, 0, 16)
	done := make(map[string]struct{}, 16)
	backtracking(0, n)
	return res
}

func backtracking(cur, n int) {
	if cur == n {
		return
	}

	tmp := make([]byte, 0, len(path))
	for _, v := range path {
		tmp = append(tmp, v)
	}

	for i, c := range tmp {

		if _, ok := done[string(path)]; ok {
			continue
		}
		backtracking(cur+1, n)
	}
}
