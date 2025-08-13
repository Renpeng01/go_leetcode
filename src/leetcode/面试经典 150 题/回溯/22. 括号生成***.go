package main

import "fmt"

// https://www.bilibili.com/video/BV14pJLzHEEE/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
// https://www.bilibili.com/video/BV1xV411q7wE/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
var res []string

// 不要想着找规律的回溯，要用暴力的思想加剪枝
func generateParenthesis(n int) []string {
	res = make([]string, 0, 256)
	backtracking("(", 1, 0, n)
	return res
}

func backtracking(s string, l, r, n int) {
	if r > l || r+l > 2*n {
		return
	}

	// fmt.Println(l, r, n)
	if l == n && r == n {
		res = append(res, s)
		// fmt.Println(1111)
		return
	}

	backtracking(s+"(", l+1, r, n)
	backtracking(s+")", l, r+1, n)
}

func main() {
	n := 3

	r := generateParenthesis(n)
	fmt.Println(r)
}
