package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	set := make(map[byte]int)
	set[s[0]] = 0
	begin := 0
	i := 1
	deleteInex := 0
	for i < len(s) {
		v, ok := set[s[i]]
		if ok {
			if (i - begin) > max {
				max = i - begin
			}
			for j := deleteInex; j <= v; j++ {
				delete(set, s[j])
			}
			deleteInex = v + 1
			begin = v + 1
		} else {
			if i == len(s)-1 && (i-begin+1) > max {
				max = i - begin + 1
			}
		}

		set[s[i]] = i
		i++

	}
	return max
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)

}
