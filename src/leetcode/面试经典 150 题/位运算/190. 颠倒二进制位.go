package main

import (
	"fmt"
	"math"
)

func reverseBits(n int) int {

	bits := make([]int, 0, 16)

	for n > 0 {
		a := n % 2
		n = n / 2
		bits = append(bits, a)
	}

	res := 0
	for i := len(bits) - 1; i >= 0; i-- {
		if bits[len(bits)-1-i] != 0 {
			res += int(math.Pow(2, float64(i)))
		}
	}
	return res
}

func main() {
	res := reverseBits(11)
	fmt.Println(res)
}
