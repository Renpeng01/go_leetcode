package main

import "fmt"

var res []string
var path []byte

func letterCombinations(digits string) []string {
	numLetters := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	res = make([]string, 0, 16)
	path = make([]byte, 0, 16)

	input := make([][]byte, 0, 16)
	for i := 0; i < len(digits); i++ {
		input = append(input, numLetters[digits[i]])
	}
	backtracking(input, 0)
	fmt.Printf("input: %+v\n", input)
	return res

}

func backtracking(input [][]byte, k int) {
	if len(path) == len(input) {
		tmp := make([]byte, 0, len(path))
		for _, c := range path {
			tmp = append(tmp, c)
		}
		res = append(res, string(tmp))
		return
	}

	for i := k; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			path = append(path, input[i][j])
			backtracking(input, k+1)
			path = path[:len(path)-1]
		}
	}
}

func main() {
	digits := "23"
	res := letterCombinations(digits)
	fmt.Printf("res: %+v\n", res)

}
