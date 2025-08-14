package main

import (
	"fmt"
	"math"
)

// 超时
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	left := 0
	right := 0
	minLen := math.MaxInt

	res := ""

	for ; len(s)-1-left+1 >= len(t); left++ {
		letterCnt := make(map[byte]int, len(t))
		for i := 0; i < len(t); i++ {
			letterCnt[t[i]]++
		}

		for i := left; i < len(s); i++ {
			if _, ok := letterCnt[s[i]]; ok {
				left = i
				break
			}
		}

		// fmt.Println("left:", left)

		isFinish := false
		for i := left; i < len(s); i++ {
			if _, ok := letterCnt[s[i]]; ok {
				letterCnt[s[i]]--

				if letterCnt[s[i]] == 0 {
					delete(letterCnt, s[i])
				}
			}
			if len(letterCnt) == 0 {
				right = i
				isFinish = true
				break
			}
		}

		// fmt.Println("right:", right, "isFinish:", isFinish)
		if left > right {
			break
		}
		// fmt.Println("res:", s[left:right+1])

		if isFinish && right+1-left < minLen {
			minLen = right + 1 - left
			res = s[left : right+1]
		}

	}
	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	// s := "aa"
	// t := "aa"

	// s := "a"
	// t := "a"

	res := minWindow(s, t)
	fmt.Println(res)
}
