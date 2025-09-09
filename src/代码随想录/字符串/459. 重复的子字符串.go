package main

func repeatedSubstringPattern(s string) bool {

	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}
		for j := 0; j < len(s); j = j + i {
			if s[:i] != s[j:j+i] {
				break
			}
			if j == len(s)-1-i {
				return true
			}
		}
	}
	return false

}
