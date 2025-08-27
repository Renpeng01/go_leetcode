package main

func isAnagram1(s string, t string) bool {
	sSet := make(map[rune]int, len(s))
	for _, c := range s {
		_, ok := sSet[c]
		if ok {
			sSet[c]++
		} else {
			sSet[c] = 1
		}
	}
	tSet := make(map[rune]int, len(t))
	for _, c := range t {
		_, ok := tSet[c]
		if ok {
			tSet[c]++
		} else {
			tSet[c] = 1
		}
	}

	if len(sSet) != len(tSet) {
		return false
	}

	for c, val := range sSet {
		v, ok := tSet[c]
		if !ok {
			return false
		}

		if v != val {
			return false
		}
	}
	return true
}

// 优化：使用一个map来实现
func isAnagram(s string, t string) bool {
	set := make(map[rune]int, len(s))
	for _, c := range s {
		_, ok := set[c]
		if ok {
			set[c]++
		} else {
			set[c] = 1
		}
	}

	for _, c := range t {
		set[c]--
	}

	for _, v := range set {
		if v != 0 {
			return false
		}
	}

	return true
}
