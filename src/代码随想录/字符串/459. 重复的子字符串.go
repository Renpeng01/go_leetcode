package main

func repeatedSubstringPattern(s string) bool {

	n := len(s)
	for step := 1; step <= n/2; step++ {
		if n%step != 0 {
			continue
		}
		for j := 0; j < len(s); j = j + step {
			if s[:step] != s[j:j+step] {
				break
			}
			if j == len(s)-1-step {
				return true
			}
		}
	}
	return false

}
