package main

func lengthOfLastWord(s string) int {
	b := -1
	s = " " + s
	for i := len(s) - 1; i >= -1; i-- {
		if b < 0 {
			if s[i] == ' ' {
				continue
			} else {
				b = i
			}
		} else {
			if i == 0 || s[i] == ' ' {
				return b - i
			}

		}

	}
	return 0

}
