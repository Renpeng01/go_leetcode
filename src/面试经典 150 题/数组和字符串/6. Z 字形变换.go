package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
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
			// fmt.Printf("11111111,i: %+v table[%+v][%+v]\n", i, i%(numRows+numRows-2), colStep)
			table[i%(numRows+numRows-2)][colStep] = s[i]
		} else {
			// fmt.Printf("22222222,i: %+v table[%+v][%+v]\n", i, numRows-1-(i%(numRows+numRows-2)-(numRows-1)), i%(numRows+numRows-2)-(numRows-1)+colStep)
			table[numRows-1-(i%(numRows+numRows-2)-(numRows-1))][i%(numRows+numRows-2)-(numRows-1)+colStep] = s[i]
		}

	}

	res := make([]byte, 0, len(s))
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[0]); j++ {
			if table[i][j] > 0 {
				res = append(res, table[i][j])
			}
		}
	}
	return string(res)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 4

	res := convert(s, numRows)
	fmt.Printf("res: %+v", res)
}
