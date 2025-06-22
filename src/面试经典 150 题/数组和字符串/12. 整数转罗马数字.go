package main

import (
	"fmt"
	"sort"
)

func intToRoman(num int) string {
	rel := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	keys := make([]int, 0, len(rel))
	for k := range rel {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := ""
	fmt.Println(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		fmt.Printf("i: %+v, key: %+v, num: %+v, num/keys[i]: %+v\n", i, keys[i], num, num/keys[i])
		if num/keys[i] > 0 {
			if num/keys[i] == 4 {
				if i == 0 {
					res += "VI"
				} else if i == 2 {
					res += "XL"
				} else if i == 4 {
					res += "CD"
				} else {
					for j := 0; j < num/keys[i]; j++ {
						res += rel[keys[i]]
					}
				}
			} else if num/keys[i] == 9 || num == 9 {
				if i == 0 || i == 1 {
					res += "IX"
					break
				} else if i == 2 {
					res += "XC"
				} else if i == 4 {
					res += "CM"
				} else {
					for j := 0; j < num/keys[i]; j++ {
						res += rel[keys[i]]
					}
				}

			} else {
				// fmt.Printf("cnt: %+v,relVal: %+v\n", num/keys[i], rel[keys[i]])
				for j := 0; j < num/keys[i]; j++ {
					res += rel[keys[i]]
				}

			}

		}
		num = num - num/keys[i]*keys[i]
	}
	return res
}

// MMMDCCXXXXIX
// MMMDCCXLIX
// MMMDCCXLIX

func main() {
	num := 1900
	res := intToRoman(num)
	fmt.Println(res)
}
