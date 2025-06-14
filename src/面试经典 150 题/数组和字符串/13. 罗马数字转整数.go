package main

func romanToInt(s string) int {
	rel := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 1 {
		return rel[s[0]]
	}

	s = " " + s
	num := 0
	currnetIdx := len(s) - 1

	for currnetIdx > 0 {
		if rel[s[currnetIdx]] <= rel[s[currnetIdx-1]] {
			num += rel[s[currnetIdx]]
		} else {
			num += rel[s[currnetIdx]] - rel[s[currnetIdx-1]]
			currnetIdx--
		}
		currnetIdx--
	}

	return num
}
