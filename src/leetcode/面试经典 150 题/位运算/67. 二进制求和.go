package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {

	i := len(a) - 1
	j := len(b) - 1

	pre := 0
	res := ""
	for i >= 0 && j >= 0 {
		av := a[i] - '0'
		bv := b[j] - '0'

		v := int(av) + int(bv) + pre
		pre = 0
		if v > 1 {
			pre = v / 2
			v = v % 2
		}
		res = strconv.Itoa(v) + res
		i--
		j--
	}
	// fmt.Println(res)

	for m := i; m >= 0; m-- {
		av := a[m] - '0'
		v := int(av) + pre
		pre = 0
		if v > 1 {
			pre = v / 2
			v = v % 2

		}
		res = strconv.Itoa(v) + res
	}
	for m := j; m >= 0; m-- {
		av := b[m] - '0'
		v := int(av) + pre
		pre = 0
		if v > 1 {
			pre = v / 2
			v = v % 2

		}
		res = strconv.Itoa(v) + res
	}

	if pre == 1 {
		res = "1" + res
	}

	return res
}

func main() {
	a := "11"
	b := "1"
	res := addBinary(a, b)
	fmt.Println(res)
}
