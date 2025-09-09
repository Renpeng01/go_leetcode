package main

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")

	left := 0
	right := len(words) - 1
	for left < right {
		words[left], words[right] = words[right], words[left]
	}
	return strings.Join(words, " ")
}
