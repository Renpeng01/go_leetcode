package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	newS := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if isEnglishLetter(s[i]) {
			newS = append(newS, s[i])
		}
	}

	fmt.Println("newS: ", string(newS))

	left := 0
	right := len(newS) - 1

	for left < right {
		if newS[left] != newS[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isEnglishLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func main() {
	s := "0P"
	res := isPalindrome(s)
	fmt.Println(res)
}
