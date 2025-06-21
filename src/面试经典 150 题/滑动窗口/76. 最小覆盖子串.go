package main

func minWindow(s string, t string) string {
	set := make(map[byte]int, len(s))
	for i := 0; i < len(t); i++ {
		set[t[i]]++
	}

	left := 0
	for i := 0; i < len(s); i++ {
		if set[s[i]] > 0 {
			set[s[i]]--
			if set[s[i]] == 0 {
				delete(set, s[i])
			}
			left = i
			break
		}
	}
	right := left
	for i := left + 1; i < len(s); i++ {
		if set[s[i]] > 0 {
			set[s[i]]--
			if set[s[i]] == 0 {
				delete(set, s[i])
			}
			if len(set) == 0 {
				right = i
				break
			}
		}
	}

	if right-left+1 <= 0 {
		return ""
	}

	for i := 0; i < len(t); i++ {
		set[t[i]]++
	}

	min := right - left + 1
	res := s[left : right+1]

	pathHas := make(map[byte]int, 8)

	for right < len(s)-1 {
		right++
		if s[right] == s[left] {
			for i := left + 1; i < right; i++ {

				if set[s[i]] > 0 {
					if pathHas[s[i]] > 0 {
						pathHas[s[i]]--
						if pathHas[s[i]] == 0 {
							delete(pathHas, s[i])
						}
						continue
					}
					left = i
					break
				}
			}

			if right-left+1 < min {
				min = right - left + 1
				res = s[left : right+1]
			}
		}
		if set[s[right]] > 0 {
			pathHas[s[right]]++
		}
	}
	return res
}

// ADOBECODEBANC
// ADOBEC
