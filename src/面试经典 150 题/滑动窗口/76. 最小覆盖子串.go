package main

import "fmt"

func minWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}
	set := make(map[byte]int, len(s))
	for i := 0; i < len(t); i++ {
		set[t[i]]++
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
		if set[s[i]] > 0 {
			path[s[i]]++
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
		if set[s[right]] > 0 {
			pathHas[s[right]]++
		}
	}
	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	res := minWindow(s, t)
	fmt.Println(res)
}
