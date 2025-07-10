package main

func canConstruct(ransomNote string, magazine string) bool {

	s := make(map[byte]int, len(magazine))

	for i := 0; i < len(magazine); i++ {
		s[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if s[ransomNote[i]] <= 0 {
			return false
		}
		s[ransomNote[i]]--

	}

	return true
}
