package main

import "strings"

func reverseWords(s string) string {
	res := make([]string, 0, 16)
	s = " " + s
	b := -1
	for i := len(s) - 1; i >= 0; i-- {
		if b < 0 {
			if s[i] == ' ' {
				continue
			}
			b = i
		} else {
			if s[i] == ' ' {
				res = append(res, s[i+1:b+1])
				b = -1
			}
		}

	}
	return strings.Join(res, " ")
}

func main() {

}
