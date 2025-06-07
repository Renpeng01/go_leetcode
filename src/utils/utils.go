package utils

import (
	"fmt"
	"strconv"
)

func PrintIntSlice(s []int) {

	t := ""
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			t += strconv.Itoa(s[i])
			continue
		}
		t += strconv.Itoa(s[i]) + ","
	}

	fmt.Println(t)
}
