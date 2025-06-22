package main

import "fmt"

func convert(s string, numRows int) string {
	col := ((len(s) / (numRows + numRows - 2)) + 1) * (numRows - 1)
	table := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		table[i] = make([]byte, col)
	}

	fmt.Printf("col: %+v\n", col)
	colStep := 0
	for i := 0; i < len(s); i++ {
		if i%(numRows+numRows-2) == 0 {
			colStep = (i / (numRows + numRows - 2)) * (numRows - 1)
		}
		if i%(numRows+numRows-2) <= numRows-1 {
			fmt.Printf("11111111,i: %+v table[%+v][%+v]\n", i, i%(numRows+numRows-2), colStep)
			table[i%(numRows+numRows-2)][colStep] = s[i]
		} else {
			fmt.Printf("22222222,i: %+v table[%+v][%+v]\n", i, numRows-1-(i%(numRows+numRows-2)-(numRows-1)), i%(numRows+numRows-2)-(numRows-1))
			table[numRows-1-(i%(numRows+numRows-2)-(numRows-1))][i%(numRows+numRows-2)-(numRows-1)] = s[i]

		}

	}

	res := ""
	for i := 0; i < len(table); i++ {
		for j := 0; i < len(table[0]); i++ {
			if string(table[i][j]) != "" {
				res += string(table[i][j])
			}
		}
	}
	return res
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3

	res := convert(s, numRows)
	fmt.Printf("res: %+v", res)
}
