package main

func convert(s string, numRows int) string {
	col := (len(s) / (numRows + numRows - 2)) + 1

	table := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		table[i] = make([]byte, col)
	}

	colStep := 0
	rowStep := 0
	for i := 0; i < len(s); i++ {

		if i/(numRows+numRows-2) <= numRows-1 {
			table[i/(numRows+numRows-2)+rowStep][colStep] = s[i]
		} else {
			table[numRows-1-i/(numRows+numRows-2)-(numRows-1)][i/(numRows+numRows-2)-(numRows-1)] = s[i]
		}

		if (i+1)%(numRows+numRows-2) == 0 {
			colStep = i * (numRows - 1)
			rowStep = i * (numRows - 1)
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
