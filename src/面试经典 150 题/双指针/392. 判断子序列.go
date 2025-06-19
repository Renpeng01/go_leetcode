package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if s == "" {
		return true
	}
	preJ := 0
	for i := 0; i < len(s); i++ {
		isFind := false
		for j := preJ; j < len(t); j++ {
			if s[i] == t[j] {
				isFind = true
				if i == len(s)-1 {
					return true
				}
				preJ = j + 1
				break
			}
		}
		if !isFind {
			return false
		}

	}
	return false
}
