package main

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)

	left := 0
	right := len(s) - 1

	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
