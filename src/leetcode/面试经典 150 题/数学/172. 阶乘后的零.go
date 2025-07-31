package main

import "fmt"

// 这个方法会因为n过大而导致total溢出
func trailingZeroes1(n int) int {

	total := 1
	for i := n; i >= 1; i-- {
		total *= i
	}

	fmt.Println(total)

	zeroCnt := 0
	for total > 0 {
		if total%10 != 0 {
			break
		}
		zeroCnt++
		total = total / 10
	}
	return zeroCnt
}
