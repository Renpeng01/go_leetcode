package main

var res []string
var path []byte

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	numLetters := [][]byte{
		[]byte{},
		[]byte{},
		[]byte{'a', 'b', 'c'},
		[]byte{'d', 'e', 'f'},
		[]byte{'g', 'h', 'i'},
		[]byte{'j', 'k', 'l'},
		[]byte{'m', 'n', 'o'},
		[]byte{'p', 'q', 'r', 's'},
		[]byte{'t', 'u', 'v'},
		[]byte{'w', 'x', 'y', 'z'},
	}
	res = make([]string, 0, 16)
	path = make([]byte, 0, 16)

	backtracking(numLetters, digits, 0)
	return res

}

func backtracking(numLetters [][]byte, digits string, startIndex int) {

	if len(path) == len(digits) {
		tmp := make([]byte, 0, len(path))
		for _, c := range path {
			tmp = append(tmp, c)
		}
		res = append(res, string(tmp))
		return
	}

	chars := numLetters[digits[startIndex]-'0']
	for _, c := range chars {
		path = append(path, c)
		backtracking(numLetters, digits, startIndex+1)
		path = path[:len(path)-1]
	}
}
