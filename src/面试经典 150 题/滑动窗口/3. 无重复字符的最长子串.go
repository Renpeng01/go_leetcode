package main

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	max := 1
	set := make(map[byte]int)
	set[s[0]] = 0

	begin := 0
	i := 1
	for i < len(s) {
		v, ok := set[s[i]]
		if ok {
			if i-begin > max {
				max = i - begin
			}
			for j := 0; j <= i; j++ {
				delete(set, s[j])
			}
			set[s[i]] = i
			i++
			begin = v + 1

		} else {
			set[s[i]] = i
			i++
		}
		if i == len(s)-1 && i-begin > max {
			max = i - begin
		}

	}
	return max
}
