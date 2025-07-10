package main

import "fmt"

// 超过限制
func lengthOfLongestSubstring1(s string) int {
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
			if _, ok := set[string(s[tail])]; !ok {
				if tail-head+1 > max {
					max = tail - head + 1
				}
			} else {
				if tail-head > max {
					max = tail - head
				}
			}
		}
		if _, ok := set[string(s[tail])]; !ok {
			set[string(s[tail])] = tail
		} else {
			if tail-head > max {
				max = tail - head
			}
			head = set[string(s[tail])] + 1

			set = make(map[string]int, len(s))
			for i := head; i <= tail; i++ {
				set[string(s[i])] = i
			}
		}
	}
	return max
}

func lengthOfLongestSubstring(s string) int {

	set := make(map[string]int)

	max := 0
	l := -1

	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(set, string(s[i-1]))
		}

		for l+1 < len(s) && set[string(s[l+1])] == 0 {
			set[string(s[l+1])]++
			l++
		}

		if l-i+1 > max {
			max = l - i + 1
		}

	}
	return max
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
