package main

import "fmt"

// 超时
// 179 / 182 个通过的测试用例
func findSubstring(s string, words []string) []int {
	res := make([]int, 0, 8)
	step := len(words[0])
	subStrLen := len(words) * step
	if len(s) < subStrLen {
		return res
	}

	wordSet := make(map[string]int, len(words))

	tempStr := ""
	for _, v := range words {
		wordSet[v]++
		tempStr += v
	}

	if len(wordSet) == 1 {

		for i := 0; i <= len(s)-subStrLen; i++ {
			if s[i:i+subStrLen] == tempStr {
				res = append(res, i)
			}
		}
		return res
	}
	// fmt.Printf("wordSet: %+v, subStrLen: %+v\n", wordSet, subStrLen)

	for i := 0; i <= len(s)-subStrLen; i++ {
		// fmt.Printf("wordSet: %+v\n", wordSet)
		// fmt.Printf("new turn i: %+v\n", i)
		revertWords := make([]string, 0, len(words))
		for b := 0; b <= subStrLen-step; b += step {
			k := s[i+b : i+b+step]
			// fmt.Printf("i: %+v b: %+v, step: %+v, (i+b): %+v,(i+b+step): %+v, key: %+v\n", i, b, step, i+b, i+b+step, k)
			if wordSet[k] > 0 {
				wordSet[k]--
				revertWords = append(revertWords, k)
				if b == subStrLen-step {
					res = append(res, i)
				}
			} else {
				break
			}
		}
		for _, w := range revertWords {
			wordSet[w]++
		}
	}
	return res
}

func main() {
	// s := "barfoothefoobarman"
	// words := []string{"foo", "bar"}
	s := "aaa"
	words := []string{"a", "a"}

	res := findSubstring(s, words)
	fmt.Printf("res: %+v\n", res)

}
