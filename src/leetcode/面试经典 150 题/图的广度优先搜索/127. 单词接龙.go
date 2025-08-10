package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	q := make([]string, 0, 32)
	vis := make(map[string]struct{}, 32)

	q = append(q, beginWord)
	vis[beginWord] = struct{}{}
	res := 1

	for len(q) > 0 {

		n := len(q)
		for i := 0; i < n; i++ {
			w := q[0]
			q = q[1:]

			if w == endWord {
				return res
			}
			for _, word := range wordList {
				_, ok := vis[word]
				if distance(w, word) == 1 && !ok {
					q = append(q, word)
					vis[word] = struct{}{}
				}
			}
		}
		res++
	}
	return 0
}

func distance(a, b string) int {
	cnt := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt
}
