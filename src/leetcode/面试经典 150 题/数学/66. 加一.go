package main

import (
	"fmt"
	"math"
)

// 这个版本当digits很长时，会导致sum溢出
func plusOne1(digits []int) []int {

	sum := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * int(math.Pow10(len(digits)-1-i))
	}
	sum++
	res := make([]int, 0, len(digits)+1)

	for sum > 0 {
		res = append(res, sum%10)
		sum = sum / 10
	}

	l := 0
	r := len(res) - 1

	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

func plusOne(digits []int) []int {
	pre := 1

	digits = append([]int{0}, digits...)
	// fmt.Println("digits:", digits)
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+pre > 9 {
			pre = (digits[i] + pre) / 10
			digits[i] = (digits[i] + pre) % 10
		} else {
			digits[i]++
			break
		}
		// fmt.Printf("digits[%+v]: %+v\n", i, digits[i])
		// fmt.Printf("pre %+v\n", pre)

	}

	if digits[0] == 0 {
		digits = digits[1:]
	}
	return digits
}

func main() {
	digits := []int{9, 9}
	res := plusOne(digits)
	fmt.Println(res)
}
