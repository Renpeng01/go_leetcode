package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 今天状态不好!!!!!!!!!!
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
		fmt.Printf("i: %+v, key: %+v, num: %+v, num/keys[i]: %+v\n", i, keys[i], num, highestDigit(num))


		if highestDigit(num) > 0 {
			if highestDigit(num) == 4 {
				if num > 100 {
					res += "CD"

				} else if num > 10 {
					res += "XL"
				} else if num > 1 {
					res += "VI"
				} else {
					for j := 0; j < num/keys[i]; j++ {
						res += rel[keys[i]]
					}
				}
				// if i == 0 {
				// 	res += "VI"
				// } else if i == 2 {
				// 	res += "XL"
				// } else if i == 4 {
				// 	res += "CD"
				// } else {
				// for j := 0; j < num/keys[i]; j++ {
				// 	res += rel[keys[i]]
				// }
				// }
			} else if highestDigit(num) == 9 || num == 9 {
				if num > 500 {
					res += "CM"
				} else if num > 50 {
					res += "XC"
				} else if num > 5 {
					res += "IX"
				} else {
					for j := 0; j < num/keys[i]; j++ {
						res += rel[keys[i]]
					}
				}
				// if i == 0 || i == 1 {
				// 	res += "IX"
				// 	break
				// } else if i == 2 {
				// 	res += "XC"
				// } else if i == 4 {
				// 	res += "CM"
				// } else {
				// for j := 0; j < num/keys[i]; j++ {
				// 	res += rel[keys[i]]
				// }
				// }

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
	num := 1994
	res := intToRoman(num)
	fmt.Println(res)
}

func highestDigit(num int) int {
	if num == 0 {
		return 0
	}

	// 处理负数情况
	if num < 0 {
		num = -num
	}

	s := strconv.Itoa(num)
	digit, _ := strconv.Atoi(string(s[0]))
	return digit
}

MCMCDXCXLVIVI
MCMXCIV