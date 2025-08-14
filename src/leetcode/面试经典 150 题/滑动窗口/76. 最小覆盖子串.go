package main

import (
	"fmt"
	"math"
)

func minWindow1(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}
	set := make(map[byte]int, len(s))
	keys := make(map[byte]struct{}, len(s))
	for i := 0; i < len(t); i++ {
		set[t[i]]++
		keys[t[i]] = struct{}{}
	}

	left := -1
	for i := 0; i < len(s); i++ {
		if set[s[i]] > 0 {
			set[s[i]]--
			if set[s[i]] == 0 {
				delete(set, s[i])
			}
			left = i
			break
		}
	}

	if left < 0 {
		return ""
	}

	right := left
	path := make(map[byte]int, 8)
	for i := left + 1; i < len(s); i++ {
		if _, ok := keys[s[i]]; ok && set[s[i]] == 0 {
			path[s[i]]++
		}
		if set[s[i]] > 0 {

			set[s[i]]--
			if set[s[i]] == 0 {
				delete(set, s[i])
			}
			if len(set) == 0 {
				right = i
				break
			}
		}
	}
	fmt.Printf("left: %+v, right: %+v\n", left, right)

	for left <= right {
		if path[s[left]] > 0 {
			path[s[left]]--
			if path[s[left]] == 0 {
				delete(path, s[left])
			}
			left++
		} else {
			break
		}
	}

	fmt.Printf("left: %+v, right: %+v\n", left, right)

	if right-left+1 <= 0 {
		return ""
	}

	for i := 0; i < len(t); i++ {
		set[t[i]]++
	}

	min := right - left + 1
	res := s[left : right+1]

	pathHas := make(map[byte]int, 8)

	for right < len(s)-1 {
		right++
		if s[right] == s[left] {
			for i := left + 1; i < right; i++ {
				if set[s[i]] > 0 {
					if pathHas[s[i]] > 0 {
						pathHas[s[i]]--
						if pathHas[s[i]] == 0 {
							delete(pathHas, s[i])
						}
						continue
					}
					left = i
					break
				}
			}

			if right-left+1 < min {
				min = right - left + 1
				res = s[left : right+1]
			}
		}

		if _, ok := keys[s[right]]; ok && set[s[right]] == 0 {
			pathHas[s[right]]++
		}
	}
	return res
}

func main() {
	// s := "ADOBECODEBANC"
	// t := "ABC"

	s := "aa"
	t := "aa"

	res := minWindow(s, t)
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	letters := make(map[byte]struct{}, len(t))
	// letterCnt := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		letters[t[i]] = struct{}{}
		// letterCnt[t[i]]++

	}

	left := -1
	right := 0
	minLen := math.MaxInt

	res := ""

	for len(s)-1-left >= len(t) {
		letterCnt := make(map[byte]int, len(t))
		for i := 0; i < len(t); i++ {
			letterCnt[t[i]]++

		}

		for i := left + 1; i < len(s); i++ {
			if _, ok := letters[s[i]]; ok {
				left = i
				break
			}
		}

		for i := left; i < len(s); i++ {
			if len(letterCnt) == 0 {
				right = i
				break
			}
			if _, ok := letters[s[i]]; ok {
				letterCnt[s[i]]--

				if letterCnt[s[i]] == 0 {
					delete(letterCnt, s[i])
				}
			}
		}

		if left >= right {
			break
		}

		if right+1-left < minLen {
			minLen = right + 1 - left
			res = s[left : right+1]
		}

	}
	return res
}
