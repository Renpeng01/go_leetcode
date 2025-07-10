package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")

	// fmt.Printf("len(str): %+v, len(s): %+v", len(strs), len(pattern))
	if len(strs) != len(pattern) {
		return false
	}
	map1 := make(map[byte]string, len(strs))
	map2 := make(map[string]byte, len(strs))
	for i := 0; i < len(strs); i++ {
		v1, ok1 := map1[pattern[i]]
		v2, ok2 := map2[strs[i]]

		// fmt.Printf("v1: %+v, ok1: %+v. v2: %+v, ok2: %+v, pattern[%+v]: %+v, strs[%+v]: %+v\n", v1, ok1, v2, ok2, i, pattern[i], i, strs[i])

		if ok1 && ok2 && v1 == strs[i] && v2 == pattern[i] {
			continue
		}

		if !ok1 && !ok2 && v1 != strs[i] && v2 != pattern[i] {
			map1[pattern[i]] = strs[i]
			map2[strs[i]] = pattern[i]
			continue
		}
		return false

	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"

	res := wordPattern(pattern, s)
	fmt.Println(res)
}
