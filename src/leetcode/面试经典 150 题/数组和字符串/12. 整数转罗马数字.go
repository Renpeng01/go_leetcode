package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 今天状态不好!!!!!!!!!!
func intToRoman1(num int) string {
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




var valueSymbols = []struct {
    value  int
    symbol string
}{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func intToRoman(num int) string {
    roman := []byte{}
    for _, vs := range valueSymbols {
        for num >= vs.value {
            num -= vs.value
            roman = append(roman, vs.symbol...)
        }
        if num == 0 {
            break
        }
    }
    return string(roman)
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/integer-to-roman/solutions/774611/zheng-shu-zhuan-luo-ma-shu-zi-by-leetcod-75rs/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。