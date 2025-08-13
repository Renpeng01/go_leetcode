package main

import "fmt"

// https://www.bilibili.com/video/BV14pJLzHEEE/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
// https://www.bilibili.com/video/BV1xV411q7wE/?spm_id_from=333.337.search-card.all.click&vd_source=70c464e99440c207e9933663bb2e5166
var exist map[string]struct{}
var res []string

func generateParenthesis(n int) []string {
	exist = make(map[string]struct{}, 256)
	res = make([]string, 0, 256)
	backtracking("(", 1, 0, n)
	return res
}

func backtracking(s string, l, r, n int) {
	if r > l || r+l > n {
		return
	}

	if l == n && r == n {
		if _, ok := exist[s]; !ok {
			res = append(res, s)
			exist[s] = struct{}{}
		}
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
