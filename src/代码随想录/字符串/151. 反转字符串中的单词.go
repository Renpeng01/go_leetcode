package main

import "strings"

func reverseWords(s string) string {

	words := make([]string, 0, 16)

	word := make([]byte, 0, 16)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(word) > 0 {
				words = append(words, string(word))
				word = make([]byte, 0, 16)
				continue
			}
			continue
		}
		word = append(word, s[i])
	}

	if len(word) > 0 {
		words = append(words, string(word))
	}

	left := 0
	right := len(words) - 1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}
