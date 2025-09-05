package main

func isAnagram1(s string, t string) bool {

	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		m[t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagram(s string, t string) bool {

	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-97]++
	}

	for i := 0; i < len(t); i++ {
		m[t[i]-97]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
