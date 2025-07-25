package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	splitStrs := make([][]string, 0, len(words))

	totalLen := len(words[0])
	tmp := make([]string, 0, 8)
	for i := 0; i < len(words); i++ {
		if i == 0 {
			tmp = append(tmp, words[0])
		} else {
			if totalLen+len(words[i])+1 > maxWidth {
				splitStrs = append(splitStrs, tmp)
				tmp = make([]string, 0, 8)
				tmp = append(tmp, words[i])
				totalLen = len(words[i])
			} else {
				tmp = append(tmp, words[i])
				totalLen += len(words[i]) + 1
			}
		}

		if i == len(words)-1 {
			splitStrs = append(splitStrs, tmp)
		}
	}

	res := make([]string, 0, len(splitStrs))
	for i, strs := range splitStrs {
		str := ""

		if len(strs) == 1 {
			str += strs[0] + strings.Repeat(" ", maxWidth-len(strs[0]))
		} else {
			wordsTotalLen := 0
			for _, clildStr := range strs {
				wordsTotalLen += len(clildStr)
			}
			paddingCnt := maxWidth - wordsTotalLen
			moreCnt := paddingCnt % (len(strs) - 1)

			if i == (len(splitStrs) - 1) {
				str = strings.Join(strs, " ") + strings.Repeat(" ", maxWidth-(len(strs)-1)-wordsTotalLen-moreCnt)
				res = append(res, str)
				continue
			}
			if moreCnt == 0 {
				str = strings.Join(strs, strings.Repeat(" ", paddingCnt/(len(strs)-1)))
			} else {
				i := 0
				for paddingCnt > 0 {
					// fmt.Println("i: ", i, " paddingCnt: ", paddingCnt, " len(strs): ", len(strs))
					strs[i%(len(strs)-1)] += " "
					paddingCnt--
					i++
				}
				str = strings.Join(strs, "")
			}
		}
		res = append(res, str)
		// fmt.Println(isLeft)
	}
	return res
}

func main() {
	// words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// words := []string{"example", "of", "text"}
	// words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	// words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	// words := []string{"Science", "is", "what", "we"}
	words := []string{"Listen", "to", "many,", "speak", "to", "a", "few."}
	res := fullJustify(words, 6)
	// fmt.Println(strings.Join(res, ","))

	for _, v := range res {
		fmt.Println(v + ",")
	}
}
