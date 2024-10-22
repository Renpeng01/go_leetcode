package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := 0
	head := 0
	tail := 1

	set := make(map[string]int, len(s))
	set[string(s[0])] = 0
	for ; tail < len(s) && head < len(s); tail++ {
		if tail == len(s)-1 {
			if tail-head > max {
				max = tail - head
			}
			break
		}
		if _, ok := set[string(s[tail])]; !ok {
			set[string(s[tail])] = tail
		} else {
			if tail-head > max {
				max = tail - head
			}
			delete(set, string(s[head]))
			set[string(s[tail])] = tail
			head = set[string(s[tail])] + 1
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
