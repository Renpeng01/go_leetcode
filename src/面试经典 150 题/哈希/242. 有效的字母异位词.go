package main

func isAnagram(s string, t string) bool {

	set := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		set[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		set[t[i]]--
	}

	for _, v := range set {
		if v != 0 {
			return false
		}
	}
	return true
}
